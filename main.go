/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-11-19 10:51
 */
package main

import (
	"fmt"
	"github.com/JunRun/Go-chatroom/rclient"
	"github.com/JunRun/Go-chatroom/server"
	"github.com/gorilla/websocket"
	"net/http"
)

func main() {
	fmt.Println("Starting application...")
	http.HandleFunc("/ws", wsHandler)
}

func wsHandler(res http.ResponseWriter, req *http.Request) {
	//将http协议升级成websocket协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}

	//每一次连接都会新开一个client，client.id通过uuid生成保证每次都是不同的
	client := &rclient.Client{Id: uuid.Must(uuid.NewV4()).String(), Socket: conn, Message: make(chan []byte)}
	//注册一个新的链接

	//启动协程收web端传过来的消息
	go client.Read()
	//启动协程把消息返回给web端
	go client.Write()
}
