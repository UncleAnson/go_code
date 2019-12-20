package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golang_code/src/chatroom/client/utils"
	"golang_code/src/chatroom/common/message"
	"net"
)

type UserProcess struct {
}

//关联一个用户登录的方法

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//fmt.Printf("userId = %d, userPwd = %s\n", userId, userPwd)
	//1.客户端dial连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建LoginMes结构体
	var LoginMes message.LoginMes
	LoginMes.UserId = userId
	LoginMes.UserPwd = userPwd
	//4.将loginMes序列化
	data, err := json.Marshal(LoginMes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)
	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}

	// 7.发送data
	//7.1 先把data的长度发送给服务器
	// 先获取data长度，转成一个表示长度的byte切片
	//7.2
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(bytes[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write fail", err)
		return
	}
	//fmt.Printf("客户端发送的长度为 %d, 内容是%s",len(data), string(data))

	// 发送消息本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write fail", err)
		return
	}
	// ~创建一个Transfer实例
	tf := utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err = ", err)
		return
	}
	// 将mes.Data反序列化为 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("反序列化失败 err =", err)
	} else if loginResMes.Code == 200 {
		//fmt.Println("登录成功")
		//2.这里还需要在客户端启动一个协程，该协程保持和服务器端的通讯
		//如果服务器有数据推送给客户端，则接收并显示在客户端的终端
		go ServerProcessMes(conn)
		//1.显示登录成功后的菜单
		for {
			ShowMenu(userId)
		}

	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
