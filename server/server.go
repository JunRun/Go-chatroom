/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-11-19 11:15
 */
package server

import (
	"encoding/json"
	"github.com/JunRun/Go-chatroom/message"
	"github.com/JunRun/Go-chatroom/rclient"
)

type ManageClient struct {
	//客户端 map 储存并管理有的长连接client，在线的为true，不在的为false
	ClientMap map[*rclient.Client]bool
	//web端发送来的的message我们用broadcast来接收，并最后分发给所有的client
	Broadcast chan []byte
	//新创建的长连接client
	Register   chan *rclient.Client
	UnRegister chan *rclient.Client
}

func NewManager() *ManageClient {
	return &ManageClient{
		ClientMap:  make(map[*rclient.Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *rclient.Client),
		UnRegister: make(chan *rclient.Client),
	}
}

func (m *ManageClient) Start() {
	for {
		//开始监听
		select {
		//如果有新连接接入，就通过channel 把链接传递给conn
		case conn := <-m.Register:
			//将链接传入管理map
			m.ClientMap[conn] = true
			message, _ := json.Marshal(&message.Message{
				Content: "A new socket has connected",
			})
			m.Send(message, conn)
		//如果有连接退出
		case conn := <-m.UnRegister:
			if _, ok := m.ClientMap[conn]; ok {
				close(conn.Message)
				delete(m.ClientMap, conn)
			}
			jsonMessage, _ := json.Marshal(&message.Message{Content: "/A socket has disconnected."})
			m.Send(jsonMessage, conn)

		case message := <-m.Broadcast:
			for conn := range m.ClientMap {
				conn.Message <- message

			}

		}
	}
}

func (m *ManageClient) Send(message []byte, ignore *rclient.Client) {
	for conn := range m.ClientMap {
		if conn != ignore {
			conn.Message <- message
		}

	}
}
