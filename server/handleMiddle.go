package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"newStudy/msq/common/message"
	"newStudy/msq/serverController/handleLogin"
)

type HandleMiddle struct {
	Conn net.Conn
	Msg  [1024]byte
}

func InitHandleMiddle(conn net.Conn){
	var msg  [1024]byte
	handle := HandleMiddle{
		Conn:conn,
		Msg:msg,
	}

	handle.Handle()
}

//用户请求分发中心
func (this *HandleMiddle)Handle(){
	//延时关闭用户连接
	defer this.Conn.Close()

	//获取用户发送消息
	for {
		n,err := this.Conn.Read(this.Msg[:1024])
		if err != nil {
			if err == io.EOF{
				fmt.Printf("%v离线\n",this.Conn.RemoteAddr(),err)
			}else {
				fmt.Printf("获取%v消息失败，断开请求 err:%v\n",this.Conn.RemoteAddr(),err)
			}
			break
		}

		//将切片反序列化为结构体
		msgStruct := message.Message{}
		err = json.Unmarshal(this.Msg[:n],&msgStruct)
		if err != nil {
			fmt.Printf("反序列化结构体失败 err:%v",err)
			break
		}

		//判断是否丢包
		if msgStruct.Size != len(msgStruct.Data) {
			fmt.Println("数据丢包，重新获取")
			break
		}

		//根据消息类型选择不同的操作
		switch msgStruct.Type {
		case message.LoginMsgType://用户登陆
			login 	:= handleLogin.InitHandleLogin(this.Conn,msgStruct.Data)
			err 	= login.Login()
		case message.RegisterMsgType://用户注册
			login 	:= handleLogin.InitHandleLogin(this.Conn,msgStruct.Data)
			err 	= login.Register()
		}

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
