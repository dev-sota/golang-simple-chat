package main

import (
	"log"
	"net/http"

	"github.com/dev-sota/golang-simple-chat/handler"
)

func main() {
	// localhost:8080 でアクセスした時に index.html を読み込む
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/chat", handler.HandleClients)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
