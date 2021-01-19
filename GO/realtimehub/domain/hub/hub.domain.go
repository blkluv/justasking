package hubdomain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"justasking/GO/realtimehub/model/hubs"
	"justasking/GO/realtimehub/model/websocketmessage"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// Get by key
func Get(hubName string, hm *hubs.HubMap) *hubs.Hub {
	hm.RLock()
	defer hm.RUnlock()

	// Does this hub exist yet
	hub, ok := hm.Hubs[hubName]

	if !ok {
		// Create a new hub
		hub = NewHub(hubName, hm)
		hm.Hubs[hubName] = hub
	}

	return hub
}

// NewHub creates and starts a hub
func NewHub(hubName string, hm *hubs.HubMap) *hubs.Hub {
	hub := &hubs.Hub{
		HubName:    hubName,
		Broadcast:  make(chan []byte),
		Register:   make(chan *hubs.Client),
		Unregister: make(chan *hubs.Client),
		Clients:    make(map[*hubs.Client]bool),
	}

	go RunQuestionBox(hm, hub)

	return hub
}

// RunQuestionBox begins the hub loop
func RunQuestionBox(hm *hubs.HubMap, h *hubs.Hub) {
	hasClients := true
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			println("[", h.HubName, "]", " A new client has registered!")
			broadCastClientCount(h.Clients)
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				println("[", h.HubName, "]", " A client has unregistered!")
				broadCastClientCount(h.Clients)
				if len(h.Clients) == 0 {
					hasClients = false
					println("[", h.HubName, "]", " No more clients!")
					break
				}
			}
		case message := <-h.Broadcast:
			println("[", h.HubName, "]", " Broadcasting: ", string(message))
			for client := range h.Clients {
				select {
				case client.Send <- []byte(message):
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}

		}
		if !hasClients {
			hm.RLock()
			defer hm.RUnlock()
			delete(hm.Hubs, h.HubName)
			break
		}
	}
}

func broadCastClientCount(clients map[*hubs.Client]bool) {
	for client := range clients {
		if client.Name == "dashboard" {
			var websocketMessage websocketmessagemodel.WebSocketMessage
			websocketMessage.MessageType = "DashboardClientCount"
			websocketMessage.MessageData = string([]byte(fmt.Sprintf(`{"ClientCount" : [%v]}`, len(clients))))
			msg, _ := json.Marshal(websocketMessage)
			client.Send <- []byte(msg)
		}
	}
}

// ReadPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func ReadPump(c *hubs.Client) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(hubs.MaxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(hubs.PongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(hubs.PongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, hubs.Newline, hubs.Space, -1))
		c.Hub.Broadcast <- message
	}
}

// WritePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func WritePump(c *hubs.Client) {
	ticker := time.NewTicker(hubs.PingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(hubs.WriteWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(hubs.Newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(hubs.WriteWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// ServeWs handles websocket requests from the peer.
func ServeWs(hubName string, allHubs *hubs.HubMap, clientName string, w http.ResponseWriter, r *http.Request) {
	conn, err := hubs.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	hub := Get(hubName, allHubs)

	client := &hubs.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256), Name: clientName}
	client.Hub.Register <- client
	go WritePump(client)
	ReadPump(client)
}

// ReceiveAndBroadcastMessage broadcasts a message to the question box hub
func ReceiveAndBroadcastMessage(hubName string, hubs *hubs.HubMap, w http.ResponseWriter, r *http.Request) {

	//hubnames will always be lower
	hubName = strings.ToLower(hubName)
	hub := Get(hubName, hubs)
	var websocketMessage websocketmessagemodel.WebSocketMessage
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&websocketMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		entryString, err := json.Marshal(websocketMessage)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			hub.Broadcast <- []byte(entryString)
		}
	}
}
