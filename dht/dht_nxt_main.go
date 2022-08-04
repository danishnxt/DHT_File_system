package main

import (
	"DHT_NXT/node"
	"fmt"
)

func main() {
	test_node := node.CreateNode("test_node_1")
	fmt.Println(test_node.GetName())
	test_node.SetPreIP("localhost")
	fmt.Println(test_node.GetPreIP())

}
