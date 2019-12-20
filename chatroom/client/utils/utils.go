package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"golang_code/src/chatroom/common/message"
	"net"
)

// 将这些方法关联到结构体中
type Transfer struct {
	// 分析应该有的字段
	// 连接
	Conn net.Conn
	// buf
	Buf [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("等待读取客户端发送的数据...")
	n, err := this.Conn.Read(this.Buf[:4])
	//fmt.Println(n) // 打印读取的byte类型长度
	if n != 4 || err != nil {
		//err = errors.New("readPkg err")
		return
	}
	fmt.Println("读到的buf =", this.Buf[:4])

	// 转换buf[:4]为uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	//根据pkgLen读取 下一次客户端发送来的 消息内容
	n, err = this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil { // int 和 uint32 不能直接比较
		fmt.Println("conn.Read err =", err)
		return
	}
	// 反序列化buf message.Message
	//var mes message.Message //在返回中已经声明
	err = json.Unmarshal(this.Buf[:pkgLen], &mes) //&!!!
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var bytes [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write fail", err)
		return
	}
	//fmt.Printf("客户端发送的长度为 %d, 内容是%s",len(data), string(data))

	// 发送消息本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write fail", err)
		return
	}
	return
}
