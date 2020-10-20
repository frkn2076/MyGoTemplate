package socket

import (
	// "app/MyGoTemplate/logger"

	// "bufio"
	// "log"
	// "net"
    // "os"
    // "net/http"
)

// var H = hub{
//     broadcast:  make(chan message),
//     register:   make(chan subscription),
//     unregister: make(chan subscription),
//     rooms:      make(map[string]map[*connection]bool),
//  }

//  type message struct {
// 	data []byte
// 	room string
// }

// var upgrader = websocket.Upgrader{
//     ReadBufferSize:  1024,
//     WriteBufferSize: 1024,
//  }

//  type connection struct {
//     ws *websocket.Conn
//     send chan []byte
//  }

//  type hub struct {
//     rooms map[string]map[*connection]bool
//     broadcast chan message
//     register chan subscription
//     unregister chan subscription
//  }

//  type subscription struct {
//     conn *connection
//     room string
//  }

// func ServeWs(w http.ResponseWriter, r *http.Request, roomId string) {
//     ws, err := upgrader.Upgrade(w, r, nil)
//     if err != nil {
//        log.Println(err.Error())
//        return
//     }
//     c := &connection{send: make(chan []byte, 256), ws: ws}
//     s := subscription{c, roomId}
//     H.register <- s
//     go s.writePump()
//     go s.readPump()
//  }

//  func (h *hub) Run() {
//     for {
//        select {
//        case s := <-h.register:
//           connections := h.rooms[s.room]
//           if connections == nil {
//              connections = make(map[*connection]bool)
//              h.rooms[s.room] = connections
//           }
//           h.rooms[s.room][s.conn] = true
//        case s := <-h.unregister:
//           connections := h.rooms[s.room]
//           if connections != nil {
//              if _, ok := connections[s.conn]; ok {
//                 delete(connections, s.conn)
//                 close(s.conn.send)
//                 if len(connections) == 0 {
//                    delete(h.rooms, s.room)
//                 }
//              }
//           }
//        case m := <-h.broadcast:
//           connections := h.rooms[m.room]
//           for c := range connections {
//              select {
//              case c.send <- m.data:
//              default:
//                 close(c.send)
//                 delete(connections, c)
//                 if len(connections) == 0 {
//                    delete(h.rooms, m.room)
//                 }
//              }
//           }
//        }
//     }
//  }

//  func (s subscription) readPump() {
//     c := s.conn
//     defer func() {
//        Hub.unregister <- s
//        c.ws.Close()
//     }()
//     c.ws.SetReadLimit(maxMessageSize)
//     c.ws.SetReadDeadline(time.Now().Add(pongWait))
//     c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
//     for {
//        _, msg, err := c.ws.ReadMessage()
//        if err != nil {
//           if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
//              log.Printf("error: %v", err)
//           }
//           break
//        }
//        m := message{msg, s.room}
//        Hub.broadcast <- m
//     }
//  }

//  func (s *subscription) writePump() {
//     c := s.conn
//     ticker := time.NewTicker(pingPeriod)
//     defer func() {
//        ticker.Stop()
//        c.ws.Close()
//     }()
//     for {
//        select {
//        case message, ok := <-c.send:
//           if !ok {
//              c.write(websocket.CloseMessage, []byte{})
//              return
//           }
//           if err := c.write(websocket.TextMessage, message); err != nil {
//              return
//           }
//        case <-ticker.C:
//           if err := c.write(websocket.PingMessage, []byte{}); err != nil {
//              return
//           }
//        }
//     }
//  }
//  func (c *connection) write(mt int, payload []byte) error {
//     c.ws.SetWriteDeadline(time.Now().Add(writeWait))
//     return c.ws.WriteMessage(mt, payload)
//  }




















// const (
//     connHost = "127.0.0.1"
//     connPort = "8080"
//     connType = "tcp"
// )

// // var M *melody.Melody = melody.New()
// var Listener net.Listener = initListener()

// func initListener() net.Listener {
//     l, err := net.Listen(connType, connHost+":"+connPort)
//     if err != nil {
// 		logger.ErrorLog("Socket listening error:", err.Error())
//         os.Exit(1)
//     }
//     return l
// }

// func Start() {
//     // l, err := net.Listen(connType, connHost+":"+connPort)
//     // if err != nil {
// 	// 	logger.ErrorLog("Socket listening error:", err.Error())
//     //     os.Exit(1)
//     // }

//     // defer Listener.Close()
//     for {
//         c, err := Listener.Accept()
//         if err != nil {
// 			logger.ErrorLog("Socket connecting error: ", err.Error())
//             return
// 		}
		
// 		logger.InfoLog("Client ", c.RemoteAddr().String(), " connected.")

//         go handleConnection(c)
//     }
// }

// func handleConnection(conn net.Conn) {
//     buffer, err := bufio.NewReader(conn).ReadBytes('\n')

//     if err != nil {
// 		logger.InfoLog("Client left.")
//         conn.Close()
//         return
//     }

//     log.Println("Client message:", string(buffer[:len(buffer)-1]))

//     conn.Write(buffer)

//     handleConnection(conn)
// }
