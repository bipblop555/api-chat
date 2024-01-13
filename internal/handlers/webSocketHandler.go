package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	host     = "localhost"
	port     = "5432"
	dbuser   = "root"
	password = "root"
	dbname   = "postgres"
)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var clients = make(map[*websocket.Conn]bool)

type WebSocketContext struct {
	C        chan os.Signal
	Listener *pq.Listener
}

func HandleWebSocketConnection(w http.ResponseWriter, r *http.Request, ctx *WebSocketContext, userId string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erreur lors de la mise à niveau du WebSocket :", err)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket connecté.")
	clients[conn] = true

	go WaitForNotification(ctx.Listener, userId, conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Erreur de lecture du WebSocket :", err)
			delete(clients, conn)
			break
		}
	}
}

func StartSQL(c chan os.Signal) {
	conninfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, dbuser, password, dbname)

	_, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	listener := pq.NewListener(conninfo, 10*time.Second, time.Minute, reportProblem)
	err = listener.Listen("new_message")
	if err != nil {
		panic(err)
	}

	ctx := &WebSocketContext{
		C:        c,
		Listener: listener,
	}

	http.HandleFunc("/ws/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		HandleWebSocketConnection(w, r, ctx, id)
	})
	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		// Extraire l'ID de l'URL
		id := strings.TrimPrefix(r.URL.Path, "/ws/")
		HandleWebSocketConnection(w, r, ctx, id)
	})

	go http.ListenAndServe(":9098", nil)

	fmt.Println("Start monitoring PostgreSQL...")
	/*	for {
		WaitForNotification(listener, id)
	}*/
}

func sendToClients(data []byte) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Erreur lors de l'envoi du message WebSocket :", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func WaitForNotification(l *pq.Listener, id string, conn *websocket.Conn) {
	for {
		select {
		case n := <-l.Notify:
			// Analysez la notification pour obtenir sender_id
			var msgData map[string]interface{}
			if err := json.Unmarshal([]byte(n.Extra), &msgData); err != nil {
				log.Println("Erreur lors de l'analyse de la notification:", err)
				continue
			}

			fmt.Printf("ID : %s\n", id)

			// Vérifiez si le sender_id correspond à l'id du WebSocket
			if senderID, ok := msgData["sender_id"].(float64); ok && strconv.Itoa(int(senderID)) == id {
				fmt.Println("Received data from channel [", n.Channel, "] :", n.Extra)
				if err := conn.WriteMessage(websocket.TextMessage, []byte(n.Extra)); err != nil {
					log.Println("Erreur lors de l'envoi du message WebSocket :", err)
					conn.Close()
					return
				}
			}

		case <-time.After(10 * time.Second):
			go func() {
				l.Ping()
			}()
		}
	}
}
