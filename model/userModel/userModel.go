package userModel

type UserModel struct {
	UserId   int    `json:"user_id"`  //用户账户
	Password string `json:"password"` //用户密码
	Nickname string `json:"nickname"` //用户昵称
}
