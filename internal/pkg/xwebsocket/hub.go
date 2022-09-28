package xwebsocket

import (
	"log"
	"sync"
)

type Hub struct {
	// Registered Clients.
	Clients map[*Client]bool

	// Inbound messages from the Clients.
	Broadcast chan []byte

	// Register requests from the Clients.
	Register chan *Client

	// Unregister requests from Clients.
	Unregister chan *Client

	Close chan int
}

var hubList = sync.Map{}

func GetOrOpenHub(hubName string) *Hub {
	hub, ok := hubList.Load(hubName)
	log.Printf("%v", ok)
	if !ok {
		hub := &Hub{
			Broadcast:  make(chan []byte),
			Register:   make(chan *Client),
			Unregister: make(chan *Client),
			Clients:    make(map[*Client]bool),
			Close:      make(chan int),
		}
		go hub.Run()
		hubList.Store(hubName, hub)
		return hub
	}
	return hub.(*Hub)
}

func GetHub(hubName string) (*Hub, bool) {
	hub, ok := hubList.Load(hubName)
	return hub.(*Hub), ok
}

func GetHubNameList() []string {
	var list []string
	hubList.Range(func(key, value any) bool {
		name, ok := key.(string)
		if ok {
			list = append(list, name)
		}
		return true
	})
	return list
}

func CloseHub(hubName string) {
	hub, ok := GetHub(hubName)
	if ok {
		hub.Close <- 1
		hubList.Delete(hubName)
	}
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		case <-h.Close:
			for client := range h.Clients {
				delete(h.Clients, client)
				close(client.Send)
			}
			log.Println("close hub")
			return
		}
	}
}
