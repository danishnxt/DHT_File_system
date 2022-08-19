package main

import (
	"DHT_NXT/node"
	"os"
)

func main() {

	port_in := ""
	ip_in := ""

	if len(os.Args) == 2 {
		port_in = os.Args[1]
	}

	if len(os.Args) == 3 {
		ip_in = os.Args[2]
	}

	// should these functions be called here or should be inside the node object?
	myNode := node.CreateNode("test_node_1", ip_in, port_in)
	go myNode.InitNodeServer()
	go myNode.MaintainRing()
	myNode.InitNodeClient() // blocking call to avoid prev funcs return prematurely
}
