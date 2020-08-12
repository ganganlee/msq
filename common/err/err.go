package err

import "errors"

var (
	USERID_NOTEXITS = errors.New("用户名不存在")
	PASSWORD_ERR    = errors.New("密码错误")
)
