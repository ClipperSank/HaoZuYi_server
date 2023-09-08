package services

import (
	"errors"
	"fmt"
	"fyno/server/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type webSocketService struct {
	messageService  models.MessageService
	upgrader        websocket.Upgrader
	connections     map[string]*websocket.Conn
	userConnections map[uuid.UUID]string
}

func NewWebSocketService(ms models.MessageService) models.WebSocketService {
	return &webSocketService{
		messageService: ms,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		connections:     make(map[string]*websocket.Conn),
		userConnections: make(map[uuid.UUID]string),
	}
}

func (ws *webSocketService) GetWebSocketID(c *gin.Context) string {
	var id string
	if cookie, err := c.Cookie("websocket-id"); err == nil {
		id = cookie
	}

	// If the WebSocket ID was not found, generate a new one
	if id == "" {
		id = fmt.Sprintf("%d", time.Now().UnixNano())
		c.SetCookie("websocket-id", id, 0, "/", "", false, true)
	}

	return id
}

func (ws *webSocketService) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	return ws.upgrader.Upgrade(w, r, responseHeader)
}

func (ws *webSocketService) AddConnection(userID uuid.UUID, id string, conn *websocket.Conn) {
	ws.connections[id] = conn
	ws.userConnections[userID] = id
}

func (ws *webSocketService) GetUserConnection(userID uuid.UUID) (*websocket.Conn, error) {
	id, ok := ws.userConnections[userID]
	if !ok {
		return nil, errors.New("user not found")
	}

	conn, ok := ws.connections[id]
	if !ok {
		return nil, errors.New("connection not found")
	}

	return conn, nil
}

func (ws *webSocketService) ListenForMessages(conn *websocket.Conn, ch chan models.Message) {
	for {
		var msg models.Message
		if err := conn.ReadJSON(&msg); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Println(err)
			}
			break
		}
		ch <- msg
	}
}

func (ws *webSocketService) HandleIncomingMessages(ch chan models.Message) {
	for {
		msg := <-ch
		// Save message to the database
		_, err := ws.messageService.CreateMessage(&msg)
		if err != nil {
			log.Println(err)
			continue
		}

		// Send message to sender and receiver
		senderConn, err := ws.GetUserConnection(msg.Sender)
		if err != nil {
			log.Println(err)
			continue
		}

		if err = ws.messageService.WriteMessage(senderConn, msg); err != nil {
			log.Println(err)
			continue
		}

		receiverConn, err := ws.GetUserConnection(msg.Receiver)
		if err != nil {
			log.Println(err)
			continue
		}

		if err = ws.messageService.WriteMessage(receiverConn, msg); err != nil {
			log.Println(err)
			continue
		}
	}
}
