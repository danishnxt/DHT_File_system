package node

// Not to self with sample function
// FUNCTION WILL NOT BE EXPORTED DUE TO START WITH lower case
// func node_test() int {
// return 1
// }

// node connection details address
type nodeAddress struct {
	IP   string
	PORT string
	HASH string
}

// set port
func (cur_node *nodeAddress) setPort(port string) {
	cur_node.PORT = port
}

// get port
func (cur_node nodeAddress) getPort() string {
	return cur_node.PORT
}

// set IP
func (cur_node *nodeAddress) setIP(IP string) {
	cur_node.IP = IP
}

// get IP
func (cur_node nodeAddress) getIP() string {
	return cur_node.IP
}

// set hash
func (cur_node *nodeAddress) setHash(hash string) {
	cur_node.HASH = hash
}

// get port
func (cur_node nodeAddress) getHash() string {
	return cur_node.HASH
}
