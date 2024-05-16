package client

import "github.com/gorilla/websocket"

type Client struct {
	conn     *websocket.Conn
	Username string
}

func CreateClient(u string, c *websocket.Conn) Client {
	return Client{Username: u, conn: c}
}

func (c *Client) CloseConnection() {
	c.conn.Close()
}

func (c *Client) ReadMessages() (int, []byte, error) {
	return c.conn.ReadMessage()
}

func (c *Client) WriteMessage(message []byte) error {
	return c.conn.WriteMessage(websocket.TextMessage, message)
}
