package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

const (
	URL = "ws://localhost:8080/ws"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(URL, nil)
	if err != nil {
		log.Fatal("Erreur de connexion:", err)
	}
	defer conn.Close()

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Déconnecté du serveur.")
				return
			}
			fmt.Printf("\r>> %s\n> ", string(msg))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			break
		}
		fmt.Print("> ")
	}
}
