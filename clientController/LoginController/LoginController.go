package LoginController

import (
	"encoding/json"
	"fmt"
	"net"
	"newStudy/msq/clientController/HandleServerMsg"
	"newStudy/msq/common/encodeMsg"
	"newStudy/msq/common/login"
	"newStudy/msq/common/message"
	"os"
)

type LoginController struct {
	login.LoginMsg
	conn net.Conn
	Content string
}

/**
用户登录
 */
func (this *LoginController)Login()  {
	//获取用户名
	fmt.Print("用户名：")
	_,err := fmt.Scan(&this.UserId)
	if err != nil {
		fmt.Printf("获取用户名失败 err:%v",err)
		return
	}

	//获取用户密码
	fmt.Print("密码：")
	_,err = fmt.Scan(&this.Password)
	if err != nil {
		fmt.Printf("获取用户密码失败 err:%v",err)
		return
	}

	//封装data
	msg := message.Message{}

	//发送消息
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Printf("服务器链接失败 err:%v\n",err)
		return
	}

	//延时关闭
	defer conn.Close()

	err = msg.Send(message.LoginMsgType,this,conn)

	//接收消息
	loginStatus,err := encodeMsg.EncodeMsg(conn)
	if err != nil{
		fmt.Printf("获取消息失败 err:%v\n",err)
		return
	}

	//将消息转为结构体
	loginRes := login.LoginResMsg{}
	err = json.Unmarshal([]byte(loginStatus),&loginRes)
	if err != nil {
		fmt.Printf("将消息转为结构体失败 err:%v",err)
		return
	}
	fmt.Println(loginRes)
	//判断登陆状态
	if loginRes.Code != 200{
		fmt.Println("登陆失败，用户名或密码错误")
		fmt.Println(loginRes)
		return
	}else {
		//登陆成功，异步接受服务器推送消息，显示登陆成功后的菜单
		go HandleServerMsg.HandleServerMsg(conn)
		this.conn = conn
		for {
			this.ShowMenu()
		}
	}
}

/**
用户注册
 */
func (this *LoginController)Register()  {
	//获取用户名
	fmt.Print("用户名：")
	_,err := fmt.Scan(&this.UserId)
	if err != nil {
		fmt.Printf("获取用户名失败 err:%v",err)
		return
	}

	//获取用户密码
	fmt.Print("密码：")
	_,err = fmt.Scan(&this.Password)
	if err != nil {
		fmt.Printf("获取用户密码失败 err:%v",err)
		return
	}

	//获取用户密码
	fmt.Print("昵称：")
	_,err = fmt.Scan(&this.Nickname)
	if err != nil {
		fmt.Printf("获取用户昵称失败 err:%v",err)
		return
	}

	//封装data
	msg := message.Message{}

	//发送消息
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Printf("服务器链接失败 err:%v\n",err)
		return
	}

	//延时关闭
	defer conn.Close()

	err = msg.Send(message.RegisterMsgType,this,conn)

	//接收消息
	loginStatus,err := encodeMsg.EncodeMsg(conn)
	if err != nil{
		fmt.Printf("获取消息失败 err:%v\n",err)
		return
	}

	//将消息转为结构体
	loginRes := login.LoginResMsg{}
	err = json.Unmarshal([]byte(loginStatus),&loginRes)
	if err != nil {
		fmt.Printf("将消息转为结构体失败 err:%v",err)
		return
	}
	//判断登陆状态
	if loginRes.Code != 200{
		fmt.Println(loginRes.Msg)
		return
	}else {
		//注册成功
		fmt.Println("注册成功，请登录。。。")
	}
}

/**
登陆操作菜单
 */
func (this *LoginController)ShowMenu(){
	fmt.Println(">>>>>>>>>>>操作菜单<<<<<<<<<<")
	fmt.Println("- 1、查看列表")
	fmt.Println("- 2、发送消息")
	fmt.Println("- 3、退出消息")
	fmt.Println("")
	fmt.Println("")
	var operation int
	_,err := fmt.Scan(&operation)
	if err != nil {
		fmt.Printf("输入错误 err:%v",err)
	}

	switch operation {
	case 1:
		fmt.Println("查看列表")
		msg := message.Message{}
		err = msg.Send(message.OnlineMsgType,this,this.conn)
		if err != nil {
			fmt.Println(err)
		}
	case 2:
		fmt.Println("input you want say...")
		_,err := fmt.Scan(&this.Content)
		if err != nil {
			fmt.Println("输入内容有误")
			break;
		}
		this.GroupChat()
	case 3:
		fmt.Println("退出系统")
		os.Exit(0)
	}
}

func (this *LoginController)GroupChat(){
	msg := message.Message{}
	err := msg.Send(message.GroupChatMsgType,this,this.conn)
	if err != nil {
		fmt.Println(err)
	}
}