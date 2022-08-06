package consts

// FILE IO CONSTS
const FILE_BUF int = 100000

// DEFAULT SERVER STUFF

const (
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
	CONN_HOST = "localhost"
)

// TCP MESSAGES

// int mapping
const (
	MSG_LIVE_PING int = 0
	MSG_DATA      int = 1
	MSG_DATA_2    int = 2
)
