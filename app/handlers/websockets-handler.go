package handlers

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
)

func WebSocketHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	clientsMu.Lock()
	clients[ws] = true
	clientsMu.Unlock()

	for {
		// Handle your business logic here
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			// log.Printf("error: %v", err)
			clientsMu.Lock()
			delete(clients, ws)
			clientsMu.Unlock()
			break
		}
		// Process the message and respond as needed
		// This is where you'd put your real-time functionality
		err = ws.WriteMessage(messageType, p)
		if err != nil {
			// log.Printf("error: %v", err)
			break
		}
	}
	return nil
}

// You can keep this function for broadcasting messages to all clients
func BroadcastMessage(message []byte) {
	clientsMu.Lock()
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			// log.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
	clientsMu.Unlock()
}
