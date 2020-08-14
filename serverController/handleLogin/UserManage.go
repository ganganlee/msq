package handleLogin

import (
	"fmt"
	"newStudy/msq/common/message"
)

//在线用户管理

var (
	UserMag *UserManage
)

type UserManage struct {
	Online map[int]*HandleLogin
}

//完成初始化操作
func init() {
	UserMag = &UserManage{
		Online: make(map[int]*HandleLogin, 1024),
	}
}

//添加修改
func (this *UserManage) AddOnlineUser(up *HandleLogin) {
	this.Online[up.UserId] = up
}

//删除
func (this *UserManage) DelOnlineUser(userId int) {
	delete(this.Online, userId)
}

//获取全部
func (this *UserManage) AllOnlineUser() map[int]*HandleLogin {
	return this.Online
}

//获取单个
func (this *UserManage) GetOnlineUserById(userId int) (up *HandleLogin,err error){
	up,ok := this.Online[userId]
	if !ok{
		return nil,fmt.Errorf("用户：%v不在线",err)
	}
	return
}

//发送上线通知
func (this *UserManage) SendOnlineNotify(user *HandleLogin) {
	list := UserMag.Online
	for key,val := range list{
		if key == user.UserId{
			continue
		}

		msg := fmt.Sprintf("用户 %v 上线",user.Nickname)
		msgStrust := message.Message{}
		msgStrust.Send(message.OnlineMsgType,msg,val.conn)
	}
}