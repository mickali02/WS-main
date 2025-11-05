// Filename: internal/ws/handler.go 
package main

 import (
	"log"
	"net/http"
	"time"
	"github.com/gorilla/webscoket"
)


 // The upgrader object is used when we need to upgrade from HTTP to RFC 6455
 var upgrader = websocket.Upgrader{}
 // Attempt to upgrade from HTTP to RFC 6455
 func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Has to be an HTTP GET request
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed",    http.StatusMethodNotAllowed) 
			return 
		}
		// Upgrade the connection from HTTP to RFC 6455
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("upgrade error: %v", err)
			return
		}
		defer conn.Close()
	// Send a control frame (high priority need immediate handling)
	_ := conn.WriteControl(
		websocket.CloseMessage,   // Close control frame (opcode = 8)
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, "bye"),
			time.Now().Add(time.Second),
		)
}
	defer conn.close()
		// Read a message
		msgType, payload, err := conn.ReadMessage()
		if err != nil {
			log.Printf("read error: %v", err)
			return
		}

		defer conn.close()
		// Process a message
		if msgType == websocket.TextMessage {
		err := conn.WriteMessage(websocket.TextMessage, payload)
		if err != nil {
			log.Printf("write error: %v", err)
			return
		}
		}
	var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
          	return true
	},


