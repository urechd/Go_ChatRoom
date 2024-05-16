package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"urechd/chatroom/client"
	"urechd/chatroom/server"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var s = server.CreateServer("Urechd's Chatroom")

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	fmt.Printf("Initiating Chatroom: %s\n", s.ChatRoomName)
	flag.Parse()
	http.HandleFunc("/chat", connectChat)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		panic(err)
	}
}

func connectChat(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	username := r.FormValue("username")

	c := client.CreateClient(username, conn)
	s.AddClient(&c)

	defer s.RemoveClient(username)

	for {
		mt, message, err := c.ReadMessages()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("MessageType: %d\n", mt)
		fmt.Printf("%s: %s\n", c.Username, message)

		err = s.BroadcastMessages(c.Username, message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
