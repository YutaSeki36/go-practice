package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:message`
}

// makeでマップを初期化　キーがウェブソケットコネクションで値がbool値
var clients = make(map[*websocket.Conn]bool)

// makeでチャンネルを初期化 チャンネルで送受信する値の型はMessage
var broadcast = make(chan Message)

// アップグレーダー　HTTPコネクションをウェブソケットにアップグレードするための構造体
var upgrader = websocket.Upgrader{}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)

	go handleMassage()

	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	// httpコネクションをウェブソケットにアップグレードする
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 関数終了時にコネクション終了
	defer ws.Close()

	// クライアントを登録
	clients[ws] = true

	// 無限ループでクライアントからのリクエストを受信
	for {
		var msg Message
		// 受け取ったJSONを構造体にマッピング
		err := ws.ReadJSON(&msg)
		log.Println(&msg)
		if err != nil {
			// エラーがあったらクライアントを登録から消す
			delete(clients, ws)
			break
		}
		// broadcastチャンネルに値を送信
		broadcast <- msg
	}
}

func handleMassage() {
	for {
		// broadcastからメッセージを受信する
		msg := <-broadcast
		log.Println("ff: ")
		log.Println(msg)

		// 登録されている全クライアントで走査
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
