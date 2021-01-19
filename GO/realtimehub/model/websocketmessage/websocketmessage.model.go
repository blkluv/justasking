package websocketmessagemodel

// WebSocketMessage is a container for passing websocket messages to the client
type WebSocketMessage struct {
	MessageType string
	MessageData string
}
