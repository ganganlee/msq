package message

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//定义消息类型常量
const (
	LoginMsgType     = "LoginMsg"
	LoginMsgResType  = "LoginResMsg"
	RegisterMsgType  = "RegisterMsg"
	OnlineMsgType    = "OnlineMsgType"
	TextMsgType      = "TextMsgType"
	GroupChatMsgType = "GroupChatMsgType" //群发消息
	ChatMsgType      = "ChatMsgType"      //私聊消息
)

//定义返回状态常量
const (
	CodeOk  = 200
	CodeErr = 400
)

//定义消息结构体
type Message struct {
	Type string //消息类型
	Data string //消息数据
	Size int    //消息长度
}

func (msg *Message) Send(msgType string, data interface{}, conn net.Conn) (err error) {
	msg.Type = msgType
	msgData, err := json.Marshal(data)
	if err != nil {
		return errors.New(fmt.Sprintf("data 序列化失败 err:%v", err))
	}
	msg.Data = string(msgData)
	msg.Size = len(msg.Data)
	str, err := json.Marshal(*msg)
	if err != nil {
		return errors.New(fmt.Sprintf("json 序列化失败 err:%v", err))
	}

	_, err = conn.Write([]byte(str))
	return err
}
