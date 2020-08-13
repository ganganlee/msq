package encodeMsg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"newStudy/msq/common/message"
)

func EncodeMsg(conn net.Conn)(data string,err error){
	msg := make([]byte,1024)
	n,err := conn.Read(msg)
	if err != nil {
		return "",errors.New(fmt.Sprintf("获取%v消息失败 err:%v",conn.RemoteAddr(),err))
	}

	//将消息转化为结构体
	msgStruct := message.Message{}
	err = json.Unmarshal(msg[:n],&msgStruct)
	if err != nil {
		return "",errors.New(fmt.Sprintf("反序列化失败 err:%v",conn.RemoteAddr(),err))
	}

	//验证数据是否丢包
	if msgStruct.Size != len(msgStruct.Data) {
		return "",errors.New("数据丢包")
	}

	//返回结果
	return msgStruct.Data,nil
}

func EncodeMsgAll(conn net.Conn)(data message.Message,err error){
	msg := make([]byte,1024)
	n,err := conn.Read(msg)
	if err != nil {
		return data,errors.New(fmt.Sprintf("获取%v消息失败 err:%v",conn.RemoteAddr(),err))
	}

	//将消息转化为结构体
	msgStruct := message.Message{}
	err = json.Unmarshal(msg[:n],&msgStruct)
	if err != nil {
		return data,errors.New(fmt.Sprintf("反序列化失败 err:%v",conn.RemoteAddr(),err))
	}

	//验证数据是否丢包
	if msgStruct.Size != len(msgStruct.Data) {
		return data,errors.New("数据丢包")
	}

	//返回结果
	return msgStruct,nil
}
