package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Conn: conn,
		Send: make(chan []byte, 256),
	}
}

// 读取消息
func (c *Client) ReadPump() {
	defer func() {
		HubInstance.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		HubInstance.Broadcast <- msg
	}
}

// 写消息
func (c *Client) WritePump() {
	defer c.Conn.Close()
	for msg := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}
