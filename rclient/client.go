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
	"github.com/JunRun/Go-chatroom/server"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id      string          //用户id
	Socket  *websocket.Conn //socket 连接
	Message chan []byte     // 发送的消息
}

func (c *Client) read() {
	defer func() {
		server.Manager.UnRegister <- c
		c.Socket.Close()
	}()
	for {
		_, info, err := c.Socket.ReadMessage()
		if err != nil {
			server.Manager.UnRegister <- c
			fmt.Println("读取信息失败，conn=", c.Id, err)
			break
		}
		jsonMe, _ := json.Marshal(&message.Message{
			Sender:  c.Id,
			Content: string(info),
		})
		server.Manager.Broadcast <- jsonMe
	}

}

func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case Info, ok := <-c.Message:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
			}
			if s := c.Socket.WriteMessage(websocket.TextMessage, Info).Error(); s {
				fmt.Println("回写信息失败", c.Id, s)
			}
		default:
			c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
		}
	}
}
