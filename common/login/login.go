package login

//定义登陆发送消息结构体
type LoginMsg struct {
	UserId int //用户账户
	Password string //用户密码
	Nickname string //用户昵称
}

//定义登陆返回消息结构体
type LoginResMsg struct {
	Code int
	Msg string
}
