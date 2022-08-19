package server

import (
	"DHT_NXT/consts"
	"DHT_NXT/util"
	"bufio"
	"fmt"
	"net"
)

func handleConnection(c net.Conn) bool {
	fmt.Println("Handle connection fired")
	conn_reader := bufio.NewReader(c)
	for {
		// create reader object on connection
		// two stage read
		net_data, err := conn_reader.ReadByte()
		print(net_data)
		util.CheckError(" HandleConnection - READ TYPE", err)

		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("conection broken - catch and resume") // graceful exit on error
				return false
			}
		}

		fmt.Println("waiting on msg -> ", net_data)
		socket_type := net_data
		// socket_type, _ := strconv.Atoi(string(net_data))
		fmt.Println("SocketType: ", socket_type)
		if socket_type == 1 {
			fmt.Println("Got ping - initial hit") // connection closed
			return true
		}
		if socket_type == 2 {
			fmt.Println("Got ping - live check") // connection closed
			return true
		}

		// IGNORE THE PAYLOAD FOR NOW //
		// data_buf := make([]byte, 10000)
		// remain_data_read, err2 := conn_reader.Read(data_buf) // read remainder
		// util.CheckError("READ PAYLOAD", err2)
		// // trim buffer
		// data_buf = data_buf[:remain_data_read] // read up to that point
		// fmt.Println(data_buf)
	}
}

func FireServer(ip string, port string) {

	l, err := net.Listen(consts.CONN_TYPE, ip+port)
	util.CheckError("INIT LISTEN ", err)
	defer l.Close()

	fmt.Println("SERVER ENGAGE " + ip + port)
	for {
		fmt.Println("GOT A CLIENT - BLOCK")
		conn, err := l.Accept()
		fmt.Println("GOT A CLIENT - UNBLOCK")
		util.CheckError("Accepting client", err)

		// This code chunk gets fired off ->
		go func() {
			// what is scope of this exit variable? Does it matter?
			exit := handleConnection(conn)
			print("Got from handle connectin => ", exit)

		}()
	}
}
