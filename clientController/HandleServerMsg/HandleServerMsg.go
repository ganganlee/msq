package HandleServerMsg

import (
	"fmt"
	"net"
)

func HandleServerMsg(conn net.Conn){
	for {
		msg 	:= make([]byte,1024)
		n,err 	:= conn.Read(msg)
		if err != nil {
			fmt.Println(err)
			break;
		}

		fmt.Printf("%v\n",string(msg[:n]))
	}
}