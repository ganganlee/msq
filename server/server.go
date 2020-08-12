package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	. "newStudy/msq/library/redisDb"
	"newStudy/msq/model/userModel"
)

var RdsPool *redis.Pool
//服务器入口文件
func main() {
	fmt.Println("服务器启中。。。")
	//端口监听
	listen,err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Printf("端口监听失败 err:%v\n",err)
	}

	//延时关闭
	defer listen.Close()

	//开启redis
	fmt.Println("redis启动中。。。")
	RdsPool = RedisPool()
	defer RdsPool.Close()

	userDao := userModel.NewUserDao(RdsPool)
	userModel.UserD = userDao

	fmt.Println("服务器启动成功，等待连接")
	//等待连接
	for {
		conn,err := listen.Accept()
		if err != nil {
			fmt.Printf("用户连接失败 err:%v\n",err)
			continue
		}

		//开辟携程处理用户请求
		go InitHandleMiddle(conn)
	}
}
