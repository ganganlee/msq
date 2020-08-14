package UserController

import (
	"fmt"
	"net"
	."newStudy/msq/clientController/LoginController"
	"os"
)

type UserController struct {
	Conn net.Conn
}

func (this *UserController)ShowMenu(login *LoginController){
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
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("退出系统")
		os.Exit(0)

	}
}