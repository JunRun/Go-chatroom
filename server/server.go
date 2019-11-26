/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-11-19 11:15
 */
package server

import (
	"github.com/JunRun/Go-chatroom/rclient"
	"github.com/gorilla/websocket"
	"net/http"
)

//web端发送来的的message我们用broadcast来接收，并最后分发给所有的client
func WsHandler(res http.ResponseWriter, req *http.Request) {
	//将http协议升级成websocket协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}

	client := rclient.NewClient(conn)
	//注册一个新的链接

	//启动协程收web端传过来的消息
	go client.Read()
	//启动协程把消息返回给web端
	go client.Write()
}
