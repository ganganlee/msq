package handleLogin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"newStudy/msq/common/login"
	"newStudy/msq/common/message"
	"newStudy/msq/model/userModel"
)

//处理用户登陆请求

type HandleLogin struct {
	conn net.Conn
	data string
	UserId int
}

func InitHandleLogin(conn net.Conn,data string) *HandleLogin {
	handle := HandleLogin{
		conn:conn,
		data:data,
	}
	return &handle
}

func (this *HandleLogin)Login()(err error){
	//将data反序列化为用户登陆结构体
	loginMsg := login.LoginMsg{}
	err = json.Unmarshal([]byte(this.data),&loginMsg)
	if err != nil {
		err = errors.New("反序列化失败")
		return
	}

	//实例化发送消息的结构体
	msgStrust := message.Message{}
	logResMsg := login.LoginResMsg{}


	userLogin := userModel.UserD
	userM,err := userLogin.GetUserInfo(loginMsg.UserId,loginMsg.Password)
	if err != nil {
		fmt.Println(err)
	}


	//比对用户名与密码
	if userM == nil {
		//用户名或密码错误
		logResMsg.Code 	= 500
		logResMsg.Msg 	= "用户名不存在"

	}else {
		//验证成功
		logResMsg.Code 	= 200
		logResMsg.Msg 	= "ok"

		//用户登陆成功，将用户信息加入到在线列表
		this.UserId = loginMsg.UserId
		UserMag.AddOnlineUser(this)

		//发送在线列表
		go this.sendOnlineUser()
	}

	//获取msgStrust json字符串
	err = msgStrust.Send(message.LoginMsgResType,logResMsg,this.conn)

	//响应数据
	return nil
}

func (this *HandleLogin)Register()(err error){
	//将data反序列化为用户登陆结构体
	loginMsg := login.LoginMsg{}
	err = json.Unmarshal([]byte(this.data),&loginMsg)
	if err != nil {
		err = errors.New("反序列化失败")
		return
	}

	//实例化发送消息的结构体
	msgStrust := message.Message{}
	logResMsg := login.LoginResMsg{}


	userLogin := userModel.UserD
	userM,err := userLogin.GetUserInfo(loginMsg.UserId,loginMsg.Password)
	if err != nil {
		fmt.Println(err)
	}

	//比对用户名与密码
	if userM != nil {
		//用户名或密码错误
		logResMsg.Code 	= 500
		logResMsg.Msg 	= "用户ID存在"
		err = msgStrust.Send(message.LoginMsgResType,logResMsg,this.conn)
		return nil;
	}

	userLogin.SetUserInfo(loginMsg.UserId,this.data)

	//返回状态
	logResMsg.Code 	= 200
	logResMsg.Msg 	= "ok"
	err = msgStrust.Send(message.LoginMsgResType,logResMsg,this.conn)

	//响应数据
	return nil
}

//发送在线列表
func (this *HandleLogin)sendOnlineUser(){
	//获取所有用户
	var list []int
	for key,_ := range UserMag.AllOnlineUser() {

		list = append(list, key)
	}

	msgStrust := message.Message{}
	msgStrust.Send(message.LoginMsgResType,logResMsg,this.conn)
	fmt.Println(list)
}