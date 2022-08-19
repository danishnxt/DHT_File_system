package node

import (
	"DHT_NXT/client"
	"DHT_NXT/consts"
	"DHT_NXT/server"
	"fmt"
	"time"
)

type Node struct {
	node_name  string
	first_node bool

	node_config_self nodeAddress
	// un-used for the moment
	node_config_pre nodeAddress
	// un-used
	node_config_post nodeAddress
}

// ================ BASIC SETTER GETTERS //

func (nodeObj *Node) GetName() string {
	return nodeObj.node_name
}
func (nodeObj *Node) SetName(nodeName string) {
	nodeObj.node_name = nodeName
}

// neighbour get/set ftns - kinda pointless?
func (nodeObj *Node) SetPreIP(new_IP string) {
	nodeObj.node_config_pre.setIP(new_IP)
}
func (nodeObj *Node) GetPreIP() string {
	return nodeObj.node_config_pre.getIP()
}
func (nodeObj *Node) SetPostIP(new_IP string) {
	nodeObj.node_config_post.setIP(new_IP)
}
func (nodeObj *Node) GetPostIP() string {
	return nodeObj.node_config_post.getIP()
}

// ================ CONSTRUCTOR AND INIT //

func CreateNode(node_name string, node_ip string, node_port string) *Node {
	first_node := false
	if node_port == "" {
		node_port = consts.CONN_PORT // default value here
		first_node = true
	}
	if node_ip == "" {
		node_ip = consts.CONN_HOST
	}
	node_obj := Node{node_name: node_name}
	node_obj.node_config_self = nodeAddress{}
	node_obj.node_config_self.setPort(node_port)
	node_obj.node_config_self.setIP(node_ip)
	node_obj.first_node = first_node

	// set neighbours to itself?
	node_obj.node_config_pre.setPort(node_port)
	node_obj.node_config_pre.setIP(node_ip)
	node_obj.node_config_post.setPort(node_port)
	node_obj.node_config_post.setIP(node_ip)

	return &node_obj // do we need this?
}

func (nodeObj *Node) InitNodeServer() {
	// I don't think we need an object type for server? Given only one per node?
	server.FireServer(nodeObj.node_config_self.getIP(), ":"+nodeObj.node_config_self.getPort())
}

func (nodeObj *Node) InitNodeClient() {
	client.FireClient()
}

func (nodeObj *Node) MaintainRing() {

	// sleep functionality here
	for {
		time.Sleep(5 * time.Second) // for now
		fmt.Println("5 SECOND PING -> ", nodeObj.first_node)
		if !nodeObj.first_node {
			// must be another node out there - find it
			if nodeObj.node_config_self.getPort() == nodeObj.node_config_post.getPort() {

				// points to self
				// assume pre also points to the same thing

				// create connection with previous
				// error check on creation of node
				fmt.Println(" ** SENDING FIRST PING ")
				alpha := client.DialNode(nodeObj.node_config_self.getIP(), consts.CONN_PORT)
				client.MessageNode(1, nil, alpha)
				client.CloseDialer(alpha)
			} else {
				// we already contacted - just do an alive check
				fmt.Println(" ** SENDING SECOND PING ")
				alpha := client.DialNode(nodeObj.node_config_self.getIP(), consts.CONN_PORT)
				client.MessageNode(2, nil, alpha)
				client.CloseDialer(alpha)
			}

		} else {
			return
		}
	}
	// infinite while loop -
	// set timer for result
}
