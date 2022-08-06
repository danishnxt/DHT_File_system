package node

import (
	"DHT_NXT/client"
	"DHT_NXT/consts"
	"DHT_NXT/server"
)

type Node struct {
	node_name  string
	first_node bool

	node_config_self nodeAddress
	// un-used for the moment
	node_config_pre  nodeAddress
	node_config_post nodeAddress
}

// ================ BASIC SETTER GETTERS //

func (nodeObj *Node) GetName() string {
	return nodeObj.node_name
}

func (nodeObj *Node) SetName(nodeName string) {
	nodeObj.node_name = nodeName
}

func (nodeObj *Node) SetPreIP(new_IP string) {
	nodeObj.node_config_pre.setIP(new_IP)
}

func (nodeObj *Node) GetPreIP() string {
	return nodeObj.node_config_pre.getIP()
}

// ================ CONSTRUCTOR AND INIT //

func CreateNode(node_name string, node_ip string, node_port string) *Node {

	first_node := false

	if node_port == "" {
		node_port = consts.CONN_PORT
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
	return &node_obj // do we need this?
}

func (nodeObj *Node) InitNodeServer() {
	// create server
	exit_chan := make(chan int)
	go server.FireServer(nodeObj.node_config_self.getIP(), ":"+nodeObj.node_config_self.getPort(), exit_chan)
	<-exit_chan // no this would never return tho
}

func (nodeObj *Node) InitNodeClient() {
	// blocking call?
	client.FireClient()
}

func (nodeObj *Node) MaintainRing() {

	// infinite while loop -
}
