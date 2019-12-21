package main

import (
	"fmt"
	"golang_code/src/chatroom/common/message"
	process2 "golang_code/src/chatroom/server/process"
	"golang_code/src/chatroom/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes message.Message) (err error) { // mes 应该调用指针
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		// 创建一个UserProcessor实例
		up := &process2.UserProcess{Conn: this.Conn}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
		up := &process2.UserProcess{Conn: this.Conn}
		err = up.ServerProcessRegister(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Processor) process2() (err error) {
	//buf := make([]byte, 8096)
	// 读客户端

	for {
		//fmt.Println("读取客户端发送的数据...")
		//n, err := conn.Read(buf[:4])
		//if n != 4 || err != nil {
		//	fmt.Println("conn.Read err=", err)
		//	return
		//}
		tf := utils.Transfer{
			Conn: this.Conn,
		}
		//fmt.Println("读到的buf=", buf[0:4])
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				fmt.Printf("客户端 %s 退出连接\n", this.Conn.RemoteAddr().String())
			} else {
				fmt.Println("readPkg err =", err)
			}
			return err
		}
		fmt.Println("接收到消息内容为：", mes)
		err = this.serverProcessMes(mes)
		if err != nil {
			fmt.Println("serverProcessMes err =", err)
			return err
		}

	}
}
