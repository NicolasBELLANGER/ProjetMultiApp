package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erreur lors de l'upgrade WebSocket", err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	log.Println("Nouvelle connexion WebSocket établie")

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Erreur WebSocket:", err)
			delete(clients, conn)
			break
		}
		log.Printf("Message reçu : %s", message)
	}
}

func notifyClients() {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte("Nouveau message reçu"))
		if err != nil {
			log.Println("Erreur lors de l'envoi du message :", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func notifyHandler(w http.ResponseWriter, r *http.Request){
	log.Println("Notification reçu de Flask")
	notifyClients()
}

func main() {
	http.HandleFunc("/ws/", handleWebSocket)
	http.HandleFunc("/notify", notifyHandler)

	log.Println("Serveur WebSocket en cours d'exécution sur le port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

