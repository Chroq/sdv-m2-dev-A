package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	PORT = ":8080"
)

type Broadcaster struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
	mutex     *sync.Mutex
	upgrader  websocket.Upgrader
}

func main() {
	broadcaster := Broadcaster{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
		mutex:     &sync.Mutex{},
		upgrader:  websocket.Upgrader{},
	}

	http.HandleFunc("/ws", broadcaster.HandleConnections)
	go broadcaster.HandleMessages()
	http.ListenAndServe(PORT, nil)
}

func (b *Broadcaster) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := b.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	b.mutex.Lock()
	b.clients[conn] = true
	b.mutex.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			b.mutex.Lock()
			delete(b.clients, conn)
			b.mutex.Unlock()
			break
		}
		b.broadcast <- msg
	}
}

func (b *Broadcaster) HandleMessages() {
	for {
		msg := <-b.broadcast
		b.mutex.Lock()
		for client := range b.clients {
			client.WriteMessage(websocket.TextMessage, msg)
		}
		b.mutex.Unlock()
	}
}
