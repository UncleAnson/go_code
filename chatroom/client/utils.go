package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golang_code/src/chatroom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	n, err := conn.Read(buf[:4])
	fmt.Println(n)
	if n != 4 || err != nil {
		//err = errors.New("readPkg err")
		return
	}
	fmt.Println("读到的buf =", buf[:4])

	// 转换buf[:4]为uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	//根据pkgLen读取 下一次客户端发送来的 消息内容
	n, err = conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil { // int 和 uint32 不能直接比较
		fmt.Println("conn.Read err =", err)
		return
	}
	// 反序列化buf message.Message
	//var mes message.Message //在返回中已经声明
	err = json.Unmarshal(buf[:pkgLen], &mes) //&!!!
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}
	return
}