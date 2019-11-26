/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-11-19 11:17
 */
package rclient

import (
	"encoding/json"
	"fmt"
	"github.com/JunRun/Go-chatroom/message"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
)

type Client struct {
	Id      string          //用户id
	Socket  *websocket.Conn //socket 连接
	Message chan []byte     // 发送的消息
}

func NewClient(socket *websocket.Conn) *Client {
	c := &Client{
		Id:      xid.New().String(),
		Socket:  socket,
		Message: nil,
	}
	AddChannel <- c
	return c
}

func (c *Client) Read() {
	defer func() {
		//发送退出信号
		QuitChannel <- c
		c.Socket.Close()
	}()
	for {
		_, info, err := c.Socket.ReadMessage()
		if err != nil {
			fmt.Println("读取信息失败，conn=", c.Id, err)
			break
		}
		jsonMe, _ := json.Marshal(&message.Message{
			Sender:  c.Id,
			Content: string(info),
		})
		message.Broadcast <- jsonMe
	}

}

func (c *Client) Write() {
	defer func() {
		QuitChannel <- c
		c.Socket.Close()
	}()

	for {
		select {
		case Info, ok := <-c.Message:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
			}
			if s := c.Socket.WriteMessage(websocket.TextMessage, Info).Error; s != nil {
				fmt.Println("回写信息失败", c.Id, s)
			}
		default:
			c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
		}
	}
}
