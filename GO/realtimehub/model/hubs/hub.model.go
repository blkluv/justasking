package hubs

import (
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// Name of this hub
	HubName string

	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

// The HubMap maintains a list of all Hubs (Boxes)
type HubMap struct {
	sync.RWMutex
	Hubs map[string]*Hub
}
