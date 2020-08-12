package LoginController

import (
	"encoding/json"
	"fmt"
	"net"
	"newStudy/msq/clientController/HandleServerMsg"
	"newStudy/msq/clientController/UserController"
	"newStudy/msq/common/login"
	"newStudy/msq/common/message"
	"newStudy/msq/common/encodeMsg"
)

type LoginController struct {
	login.LoginMsg
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
		go HandleServerMsg.HandleServerMsg()

		user := UserController.UserController{
			Conn:conn,
		}

		for {
			user.ShowMenu()
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
	fmt.Println(loginRes)
	//判断登陆状态
	if loginRes.Code != 200{
		fmt.Println("登陆失败，用户名或密码错误")
		fmt.Println(loginRes)
		return
	}else {
		//登陆成功，异步接受服务器推送消息，显示登陆成功后的菜单
		go HandleServerMsg.HandleServerMsg()

		user := UserController.UserController{
			Conn:conn,
		}

		for {
			user.ShowMenu()
		}
	}
}