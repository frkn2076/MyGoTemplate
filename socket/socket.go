package socket

import(
	"app/MyGoTemplate/logger"

	"bufio"
    "log"
    "net"
    "os"
)

const (
    connHost = "127.0.0.1"
    connPort = "8081"
    connType = "tcp"
)

// var M *melody.Melody = melody.New()
var Listener net.Listener = initListener()

func initListener() net.Listener {
    l, err := net.Listen(connType, connHost+":"+connPort)
    if err != nil {
		logger.ErrorLog("Socket listening error:", err.Error())
        os.Exit(1)
    }
    return l
}

func Start() {
    // fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
    // l, err := net.Listen(connType, connHost+":"+connPort)
    // if err != nil {
	// 	logger.ErrorLog("Socket listening error:", err.Error())
    //     os.Exit(1)
    // }

    for {
        c, err := Listener.Accept()
        if err != nil {
			logger.ErrorLog("Socket connecting error: ", err.Error())
            return
		}
		
		logger.InfoLog("Client ", c.RemoteAddr().String(), " connected.")

        go handleConnection(c)
    }
}

func handleConnection(conn net.Conn) {
    buffer, err := bufio.NewReader(conn).ReadBytes('\n')

    if err != nil {
		logger.InfoLog("Client left.")
        conn.Close()
        return
    }

    log.Println("Client message:", string(buffer[:len(buffer)-1]))

    conn.Write(buffer)

    handleConnection(conn)
}
