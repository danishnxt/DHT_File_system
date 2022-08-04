package server

import (
	"DHT_NXT/consts"
	"DHT_NXT/util"
	"bufio"
	"fmt"
	"net"
	"strconv"
	"sync"
)

func handleConnection(c net.Conn) bool {
	fmt.Println("Handle connection fired")
	for {
		// create reader object on connection
		// two stage read
		conn_reader := bufio.NewReader(c)
		net_data, err := conn_reader.ReadString('@') // delim
		util.CheckError("READ TYPE", err)
		if err.Error() == "EOF" {
			fmt.Println("conection broken - catch and resume") // graceful exit on error
			return false
		}
		net_data = net_data[:1] // remove delim
		socket_type, _ := strconv.Atoi(net_data)
		fmt.Println("SocketType: ", socket_type)
		if socket_type == 1 {
			fmt.Println("Got ping - close server") // connection closed
			return true
		}
		if socket_type == 2 {
			fmt.Println("Got ping - disconnnect client") // connection closed
			return false
		}
		data_buf := make([]byte, 10000)
		remain_data_read, err2 := conn_reader.Read(data_buf) // read remainder
		util.CheckError("READ PAYLOAD", err2)
		// trim buffer
		data_buf = data_buf[:remain_data_read] // read up to that point
		fmt.Println(data_buf)
	}
}

func FireServer(port string, exit_chan chan int) {

	var mu sync.Mutex
	var global_exit bool = false

	l, err := net.Listen(consts.CONN_TYPE, consts.CONN_HOST+port)
	util.CheckError("INIT LISTEN ", err)
	defer l.Close()

	fmt.Println("SERVER ENGAGE " + consts.CONN_HOST + port)
	for {
		conn, err := l.Accept()
		util.CheckError("Accepting client", err)
		go func() {
			// what is the scope of this exit variable?
			exit := handleConnection(conn) // I don't think this should work?
			if !exit {
				mu.Lock()
				global_exit = true
				mu.Unlock()
			}
		}()
		mu.Lock()
		if global_exit {
			mu.Unlock()
			break
		}
	}
	exit_chan <- 1
}
