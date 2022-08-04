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
