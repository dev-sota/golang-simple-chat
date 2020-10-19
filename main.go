package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 接続されるクライアント（WebSocketのコネクションに対するポインタを定義）
var clients = make(map[*websocket.Conn]bool)

// クライアントから受け取るメッセージを格納
var broadcast = make(chan Message)

// WebSocketの更新
var upgrader = websocket.Upgrader{}

// クライアントからは JSON 形式で受け取る
type Message struct {
	Message string
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// websocket の状態を更新
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket::", err)
	}
	// websocket を閉じる
	defer ws.Close()

	// 受け取ったリクエストをクライアントとして登録
	clients[ws] = true

	// WebSocketからのメッセージを待ち続ける
	for {
		var msg Message
		// メッセージ読み込み
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error occurred while reading message: %v", err)
			delete(clients, ws)
			break
		}
		// メッセージを受け取ったらbroadcastチャネルに送る
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// メッセージ受け取り
		msg := <-broadcast
		// 現在接続しているクライアント全てにメッセージを送信する
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error occurred while writing message to client: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	// localhost:8080 でアクセスした時に index.html を読み込む
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/chat", handleConnections)
	go handleMessages()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
