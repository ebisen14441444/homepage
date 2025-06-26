package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type CheckUpdate struct {
	Key   string `json:"key"`
	Value bool   `json:"value"`
}

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan CheckUpdate)
	statusMap = make(map[string]bool)
	mu        sync.Mutex
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // CORS許可
	}
)

func main() {
	e := echo.New()

	// 静的ファイル（例: static/index.html, script.js など）を配信
	e.Static("/", "static")

	// WebSocketエンドポイント
	e.GET("/ws", func(c echo.Context) error {
		return handleConnections(c.Response(), c.Request())
	})

	// WebSocket送信用ループ
	go handleBroadcast()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server started at :" + port)
	returnErr := e.Start(":" + port)
	if returnErr != nil {
		log.Fatal(returnErr)
	}
}


func handleConnections(w http.ResponseWriter, r *http.Request) error{
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return err
	}
	defer ws.Close()

	clients[ws] = true

	// 初回に状態全部送る
	mu.Lock()
	for k, v := range statusMap {
		ws.WriteJSON(CheckUpdate{Key: k, Value: v})
	}
	mu.Unlock()

	for {
		var update CheckUpdate
		if err := ws.ReadJSON(&update); err != nil {
			log.Println("read error:", err)
			delete(clients, ws)
			break
		}

		mu.Lock()
		statusMap[update.Key] = update.Value
		mu.Unlock()

		broadcast <- update
	}
	return nil
}

func handleBroadcast() {
	for update := range broadcast {
		for client := range clients {
			err := client.WriteJSON(update)
			if err != nil {
				log.Println("write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
