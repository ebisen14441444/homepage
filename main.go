package main

import (
	"log"
	"net/http"
	"sync"
	"os"

	"github.com/gorilla/websocket"
)

type CheckUpdate struct {
	Key   string `json:"key"`   // 例: "1-pink"
	Value bool   `json:"value"` // true or false
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
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)
	go handleBroadcast()

	log.Println("Server started at :8080")
	// ✅ 変更後（Render対応）
port := os.Getenv("PORT")
if port == "" {
    port = "8080" // ローカル開発用のデフォルト
}
log.Fatal(http.ListenAndServe(":"+port, nil))

}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
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
