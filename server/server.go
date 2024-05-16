package server

import (
	"fmt"
	"urechd/chatroom/client"
)

type Server struct {
	clients      map[string]*client.Client
	ChatRoomName string
}

func CreateServer(name string) Server {
	return Server{make(map[string]*client.Client), name}
}

func (s *Server) AddClient(c *client.Client) {
	fmt.Printf("%s connected to the server\n", c.Username)
	s.clients[c.Username] = c
}

func (s *Server) RemoveClient(username string) {
	s.clients[username].CloseConnection()
	fmt.Printf("%s left the server\n", username)
	delete(s.clients, username)
}

func (s *Server) BroadcastMessages(username string, message []byte) error {
	for u, c := range s.clients {
		if u != username {
			if err := c.WriteMessage(message); err != nil {
				return err
			}
		}
	}

	return nil
}
