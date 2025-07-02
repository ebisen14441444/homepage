package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"checkapp/handler"

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
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func main() {
	e := echo.New()

	// 静的ファイルを配信
	e.Static("/", "static")

	e.GET("/memo", handler.GetMemos)
	e.POST("/memo", handler.CreateMemo)
	e.DELETE("/memo/:id", handler.DeleteMemos)
	// WebSocketエンドポイント
	e.GET("/ws", func(c echo.Context) error {
		return handleConnections(c.Response(), c.Request())
	})
	// ブロードキャスト処理開始
	go handleBroadcast()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server started at :" + port)
	log.Fatal(e.Start(":" + port))
}

func handleConnections(w http.ResponseWriter, r *http.Request) error {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return err
	}
	defer ws.Close()

	clients[ws] = true

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
			if err := client.WriteJSON(update); err != nil {
				log.Println("write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
