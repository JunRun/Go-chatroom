/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-11-19 11:15
 */
package server

import "github.com/JunRun/Go-chatroom/rclient"

type ManageClient struct {
	clientMap map[*rclient.Client]bool
}
