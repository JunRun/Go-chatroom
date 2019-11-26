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
	"net/http"
)

func main() {
	fmt.Println("Starting application...")
	http.HandleFunc("/ws", server.WsHandler)
	mange := rclient.NewManager()
	mange.Server()

}
