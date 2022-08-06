package main

import (
	"DHT_NXT/node"
	"os"
)

func main() {

	port_in := ""
	ip_in := ""

	if len(os.Args) == 2 {
		port_in = os.Args[2]
	}

	if len(os.Args) == 3 {
		ip_in = os.Args[3]
	}

	myNode := node.CreateNode("test_node_1", ip_in, port_in)
	myNode.InitNodeServer()
	myNode.MaintainRing()
	myNode.InitNodeClient() // blocking call
}
