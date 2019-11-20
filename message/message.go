/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-11-19 13:37
 */
package message

//消息类型
type Message struct {
	Sender    string `json:"sender"`    //发送者
	Recipient string `json:"recipient"` //接收者
	Content   string `json:"content"`
}
