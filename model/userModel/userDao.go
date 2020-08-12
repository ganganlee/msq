package userModel

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var UserD *UserDao

type UserDao struct {
	rds *redis.Pool
}

func NewUserDao(rds *redis.Pool)*UserDao{
	dao := &UserDao{
		rds:rds,
	}
	return dao
}

func (this *UserDao)GetUserInfo(userId int,password string)(user *UserModel,err error){
	//从缓存查找数据
	redisConn := this.rds.Get()
	res ,err := redis.Bytes(redisConn.Do("hget","msq_user",userId))
	if err != nil {
		return nil,err
	}

	model := UserModel{}
	err = json.Unmarshal(res,&model)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	return &model,nil
}

/**
保存用户信息
 */
func (this *UserDao)SetUserInfo(userId int,data string)(user *UserModel,err error){
	redisConn := this.rds.Get()
	redisConn.Do("hset","msq_user",userId,data)

	return nil,nil
}
