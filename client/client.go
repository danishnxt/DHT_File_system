package client

import (
	"DHT_NXT/util"
	"bufio"
	"fmt"
	"net"
	"os"
)

func FireClient() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" ENTER PORT TO ACCESS >> ")
	port_text, _ := reader.ReadString('\n')
	// TODO: will need input validation here
	fmt.Println("Connect to server port -> ", port_text)

	CONNECT := "localhost:" + port_text
	c, err := net.Dial("tcp", CONNECT)
	util.CheckError("DIALING FROM CLIENT", err)

	fmt.Println(" ENTER MESSAGE TO SEND >> ")
	text, _ := reader.ReadString('\n')
	fmt.Fprintf(c, text+"\n")
}

// ==== CONNECTING/INTERACTING TO Nodes as client

func DialNode(IP string, port string) net.Conn {
	CONNECT := IP + ":" + port
	c, err := net.Dial("tcp", CONNECT)
	util.CheckError("DialNode function", err)
	return c
}

func MessageNode(msg_type int, data []byte, dialer_conn net.Conn) {
	send_buf := util.BuildBuffer(msg_type, data)
	fmt.Print(send_buf)
	_, err := dialer_conn.Write(send_buf)
	util.CheckError("Sending message to Node", err)
}

func CloseDialer(c net.Conn) {
	err := c.Close()
	util.CheckError("Close Dialer", err)
}
