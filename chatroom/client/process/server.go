package process

import (
	"fmt"
	"golang_code/src/chatroom/client/utils"
	"golang_code/src/chatroom/common/message"
	"net"
	"os"
)

// 显示登录成功后的界面
func ShowMenu(userId int) {
	fmt.Printf("-------------恭喜[%v]登录成功-------------\n", userId)
	fmt.Println("-------------1.显示在线用户列表-----------")
	fmt.Println("-------------2.发送消息-------------")
	fmt.Println("-------------3.信息列表-------------")
	fmt.Println("-------------4.退出系统-------------")
	fmt.Println("请选择（1-4）：")
	var key int
	fmt.Scanln(&key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表-")
	case 2:
		fmt.Println("发送消息-")
	case 3:
		fmt.Println("信息列表-")
	case 4:
		fmt.Println("选择退出-")
		os.Exit(0)

	default:
		fmt.Println("输入的选项不正确-")

	}
}

func ServerProcessMes(conn net.Conn) {
	// 创建Transfer实例，不断地读取服务器发送的消息
	tf := utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Printf("[协程]客户端 %s 正在读取服务器发送的消息\n", conn.RemoteAddr().String())
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err =", err)
			return
		}
		fmt.Println("读取消息 mes =", mes)

		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//处理
			//1.取出NotifyUserStatusMes
			//2.保存到客户端本地的map中
		default:
			fmt.Println("服务器端返回了未知类型")
		}
	}

}
