package main

import (
	"fmt"
	"net"
	"newStudy/msq/clientController/LoginController"
	"os"
)

//定义全局用户名跟密码
var userId int
var passWd string
var userName string

var conn net.Conn
func main() {
	var operation int
	for{
		basicView()
		_,err := fmt.Scan(&operation)
		if err != nil{
			fmt.Printf("获取用户命令失败 err:%v",err)
			return
		}

		//判断用户操作
		switch operation {
		case 1://用户登录、注册
			login := LoginController.LoginController{}
			login.Login()
		case 2:
			login := LoginController.LoginController{}
			login.Register()
		case 3://退出应用
			os.Exit(0)
		default:
			fmt.Println("输入指令错误，请重新输入。。。\n")
		}
	}
}

/**
显示基础操作视图
*/
func basicView(){
	fmt.Println("----------欢迎登陆messageQ聊天系统----------")
	fmt.Println("1、用户登陆")
	fmt.Println("2、用户注册")
	fmt.Println("3、关闭系统")
	fmt.Println("请输入1-3进行操作\n")
	fmt.Print("操作：")
}