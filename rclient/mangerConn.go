/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-11-20 14:17
 */
package rclient

import (
	"fmt"
	"sync"
)

var AddChannel = make(chan *Client, 10)
var QuitChannel = make(chan *Client, 10)

type ManageClient struct {
	//客户端 map 储存并管理有的长连接client，在线的为true，不在的为false
	ClientMap map[*Client]bool
	Lock      sync.RWMutex
}

func NewManager() *ManageClient {
	return &ManageClient{
		ClientMap: make(map[*Client]bool),
	}
}

//添加连接
func (m *ManageClient) AddConn(client *Client) {
	m.Lock.Lock()
	m.ClientMap[client] = true
	m.Lock.Unlock()
	fmt.Println("a new socket add", client.Id)
}

//移除连接
func (m *ManageClient) RemoveConn(client *Client) {
	m.Lock.Lock()
	delete(m.ClientMap, client)
	m.Lock.Unlock()
	fmt.Println("a socket quit", client.Id)
}

func (m *ManageClient) Server() {
	for {
		select {
		case conn := <-AddChannel:
			m.AddConn(conn)
		case conn := <-QuitChannel:
			m.RemoveConn(conn)
		}
	}
}
