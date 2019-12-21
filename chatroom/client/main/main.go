package main

import (
	"fmt"
	"golang_code/src/chatroom/client/process"
)

// 定义全局变量，因为很多函数需要用到
var userId int
var userPwd string
var userPwd2 string
var userName string

func main() {
	var key int
	var loop = true

	for loop {
		fmt.Println("---------------欢迎登录多人聊天系统--------------")
		fmt.Println("\t\t1 登录聊天室")
		fmt.Println("\t\t2 注册用户")
		fmt.Println("\t\t3 退出系统")
		fmt.Println("\t\t请选择（1-3）：")

		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入用户的ID号：")
			//fmt.Scanf("%d\n", &userId)
			fmt.Scanln(&userId)
			fmt.Println("请输入用户的密码：")
			//fmt.Scanf("%s\n", &userPwd)
			fmt.Scanln(&userPwd)
			fmt.Println(userId, userPwd)
			//loop = false
			// 完成登录
			//1.创建一个UserProcess实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd) // 虽然有返回err，但是不需要处理，因为在Login内已经有相关的处理（打印）

		case 2:
			//loop = false
			fmt.Println("请输入用户的ID号：")
			//fmt.Scanf("%d\n", &userId)
			fmt.Scanln(&userId)
			fmt.Println("请输入用户的昵称：")
			fmt.Scanln(&userName)

			for {
				fmt.Println("请输入用户的密码：")
				//fmt.Scanf("%s\n", &userPwd)
				fmt.Scanln(&userPwd)
				fmt.Println("请再次输入密码：")
				//fmt.Scanf("%s\n", &userPwd)
				fmt.Scanln(&userPwd2)
				if userPwd == userPwd2 {
					// 两次密码相同，处理注册请求
					up := &process.UserProcess{}
					up.Register(userId, userPwd, userName) // 虽然有返回err，但是不需要处理，因为在Login内已经有相关的处理（打印）

				} else {
					fmt.Println("两次密码不相同，请重新输入")
				}
			}
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	// 根据用户输入显示新菜单
	//if key == 1 {
	//	fmt.Println("请输入用户的ID号：")
	//	//fmt.Scanf("%d\n", &userId)
	//	fmt.Scanln(&userId)
	//	fmt.Println("请输入用户的密码：")
	//	//fmt.Scanf("%s\n", &userPwd)
	//	fmt.Scanln(&userPwd)
	//	fmt.Println(userId, userPwd)
	//	// 登录
	//	// 因为使用了分层模式（MVC），不能直接在这里调login函数接口
	//	//login(userId, userPwd)
	//	//if err != nil {
	//	//	fmt.Println("登录失败")
	//	//} else {
	//	//	fmt.Println("登录成功")
	//	//}
	//} else if key == 2 {
	//	fmt.Println("注册用户")
	//}
	//time.Sleep(time.Second*10)
}
