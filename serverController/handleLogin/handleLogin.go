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
	Nickname string
	Content string
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
		this.UserId 	= loginMsg.UserId
		this.Nickname 	= userM.Nickname
		UserMag.AddOnlineUser(this)

		//发送在线列表
		go this.SendOnlineUser()
		//发送用户上线
		go UserMag.SendOnlineNotify(this)

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
func (this *HandleLogin)SendOnlineUser(){
	//获取所有用户
	list := UserMag.AllOnlineUser()

	//过滤自己
	aa := UserManage{
		Online: make(map[int]*HandleLogin, 1024),
	}
	for key,val := range list{
		if key != this.UserId {
			aa.Online[key] = val
		}
	}

	msgStrust := message.Message{}
	msgStrust.Send(message.OnlineMsgType,aa.Online,this.conn)
}

//初始化用户信息
func (this *HandleLogin)EncodeUserInfo()error{
	//将data反序列化为用户登陆结构体
	loginMsg := login.LoginMsg{}
	err := json.Unmarshal([]byte(this.data),&loginMsg)
	if err != nil {
		return errors.New("反序列化失败")

	}

	this.UserId = loginMsg.UserId
	userLogin := userModel.UserD
	userM,err := userLogin.GetUserInfo(loginMsg.UserId,loginMsg.Password)
	if err != nil {
		return  err
	}
	this.Nickname 	= userM.Nickname
	this.Content	= loginMsg.Content
	return nil
}

//获取在线列表
func (this *HandleLogin)GetOnlineList(){
	err := this.EncodeUserInfo()
	if err != nil {
		msg := fmt.Sprintf("%v",err)
		msgStrust := message.Message{}
		msgStrust.Send(message.TextMsgType,msg,this.conn)
	}

	this.SendOnlineUser()
}

func (this *HandleLogin)GroupChat(){


	err := this.EncodeUserInfo()
	if err != nil {
		msg := fmt.Sprintf("%v",err)
		msgStrust := message.Message{}
		msgStrust.Send(message.TextMsgType,msg,this.conn)
	}

	//遍历在线列表，群发消息
	list := UserMag.AllOnlineUser()

	msgStrust := message.Message{}
	for key,val := range list{
		if key != this.UserId {
			msg := fmt.Sprintf("%v:%v",this.Nickname,this.Content)
			msgStrust.Send(message.TextMsgType,msg,val.conn)
		}
	}
}