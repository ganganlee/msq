package HandleServerMsg

import (
	"encoding/json"
	"fmt"
	"net"
	"newStudy/msq/common/encodeMsg"
	"newStudy/msq/common/message"
	"newStudy/msq/serverController/handleLogin"
)

func HandleServerMsg(conn net.Conn){
	for {
		msg,err := encodeMsg.EncodeMsgAll(conn)
		if err != nil{
			fmt.Println(err)
			break
		}

		//判断消息类型
		switch msg.Type {
		case message.OnlineMsgType:
			HandleOnlineMsg(msg.Data)
		}
	}
}


func HandleOnlineMsg(data string){

	online := handleLogin.UserManage{}
	err := json.Unmarshal([]byte(data),&online.Online)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("++++++++++在线列表++++++++++")
	list := online.Online
	for key,val := range list{
		fmt.Printf("ID:%v 昵称:%v\n",key,val.Nickname)
	}
}