package main

import (
	"fmt"
	"net"
)

//func serverProcessLogin(conn net.Conn, mes message.Message) (err error) {
//	// 从mes中取出mes.Data，并直接反序列化成LoginMes
//	var loginMes message.LoginMes
//	err = json.Unmarshal([]byte(mes.Data), &loginMes)
//	if err != nil {
//		fmt.Println("json.Marshal err =", err)
//		return
//	}
//	// 1.声明一个resMes
//	var resMes message.Message
//	resMes.Type = message.LoginResMesType
//	//2.声明一个LoginResMes，并完成赋值
//	var loginResMes message.LoginResMes
//	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
//		//合法
//		loginResMes.Code = 200
//	} else {
//		//不合法
//		//fmt.Println("用户名 密码 不合法")
//		loginResMes.Code = 500
//		loginResMes.Error = "请注册再使用"
//
//	}
//	//3.序列化LoginResMes
//	data, err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Println("json.Marshal err =", err)
//		return
//	}
//	//4.将data赋值给resMes的Data
//	resMes.Data = string(data) // string resMes.Data []byte data
//	data, err = json.Marshal(resMes)
//	if err != nil {
//		fmt.Println("json.Marshal err =", err)
//		return
//	}
//	// 5.发送数据
//	err = writePkg(conn, data)
//	return
//}

//func writePkg(conn net.Conn, data []byte) (err error) {
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	var bytes [4]byte
//	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
//	// 发送长度
//	n, err := conn.Write(bytes[0:4])
//	if n != 4 || err != nil {
//		fmt.Println("conn.Write fail", err)
//		return
//	}
//	//fmt.Printf("客户端发送的长度为 %d, 内容是%s",len(data), string(data))
//
//	// 发送消息本身
//	n, err = conn.Write(data)
//	if n != int(pkgLen) || err != nil {
//		fmt.Println("conn.Write fail", err)
//		return
//	}
//	return
//}

//func serverProcessMes(conn net.Conn, mes message.Message) (err error) { // mes 应该调用指针
//	switch mes.Type {
//	case message.LoginMesType:
//		// 处理登录
//		err = serverProcessLogin(conn, mes)
//		return
//	case message.RegisterMesType:
//	// 处理注册
//	default:
//		fmt.Println("消息类型不存在，无法处理")
//	}
//	return
//}

//func readPkg(conn net.Conn) (mes message.Message, err error) {
//	buf := make([]byte, 8096)
//	fmt.Println("读取客户端发送的数据...")
//	n, err := conn.Read(buf[:4])
//	//fmt.Println(n) // 打印读取的byte类型长度
//	if n != 4 || err != nil {
//		//err = errors.New("readPkg err")
//		return
//	}
//	fmt.Println("读到的buf =", buf[:4])
//
//	// 转换buf[:4]为uint32类型
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[:4])
//
//	//根据pkgLen读取 下一次客户端发送来的 消息内容
//	n, err = conn.Read(buf[:pkgLen])
//	if n != int(pkgLen) || err != nil { // int 和 uint32 不能直接比较
//		fmt.Println("conn.Read err =", err)
//		return
//	}
//	// 反序列化buf message.Message
//	//var mes message.Message //在返回中已经声明
//	err = json.Unmarshal(buf[:pkgLen], &mes) //&!!!
//	if err != nil {
//		fmt.Println("json.Unmarshal err =", err)
//		return
//	}
//	return
//}

func process(conn net.Conn) {
	// 调用总控，创建
	processor := &Processor{
		Conn: conn,
	}
	//延时关闭
	defer conn.Close()
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端通讯协程错误err =", err)
		return
	}

}

func main() {
	// 提示
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	// listen也需要延时关闭
	defer listen.Close()
	// 监听成功
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("list.Accept err=", err)
		}
		// 连接成功，启动协程
		go process(conn)
	}
}
