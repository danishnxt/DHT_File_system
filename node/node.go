package node

import (
	"DHT_NXT/client"
	"DHT_NXT/consts"
	"DHT_NXT/server"
)

type Node struct {
	node_name        string
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

func CreateNode(node_name string) *Node {
	node_obj := Node{node_name: node_name}
	node_obj.node_config_self = nodeAddress{}
	return &node_obj // do we need this?
}

func InitNode(port string) {

	if port == "" {
		port = consts.CONN_PORT
	}

	// create server
	exit_chan := make(chan int)
	go server.FireServer(":"+port, exit_chan)

	// create client
	client.FireClient()

	<-exit_chan

}

// although probably these function could directly access those of the node_address too?
// since they are the same package it should not require to be exported?
