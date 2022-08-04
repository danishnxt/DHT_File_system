package node

// ============== TESTING STUFF //

// NOTE TO SELF SAMPLE FUNCTION
// FUNCTION WILL BE EXPORTED DUE TO START WITH CAPITAL LETTER
// func NodeTest2() int {
// return node_test() + 1
// }

// func NodeTestExport() {
// 	address := nodeAddress{}
// 	address.setPort("1234")
// 	println(address.getPort())
// }

type Node struct {
	node_name        string
	node_config_self nodeAddress
	// un-used for the moment
	node_config_pre  nodeAddress
	node_config_post nodeAddress
}

// constructor
func CreateNode(node_name string) *Node {
	node_obj := Node{node_name: node_name}
	node_obj.node_config_self = nodeAddress{}
	return &node_obj
}

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

// although probably these function could directly access those of the node_address too?
// since they are the same package it should not require to be exported?
