package process2

import (
	"encoding/json"
	"fmt"
	"golang_code/src/chatroom/common/message"
	"golang_code/src/chatroom/server/model"
	"golang_code/src/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		// 对在线用户进行遍历推送
		if id == userId {
			continue
		}
		//开始通知
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	// 组装消息实体
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline
	// 将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)

	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	//发送
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println()
		return
	}
}

// 专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes message.Message) (err error) {
	// 从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	// 1.声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//2.声明一个LoginResMes，并完成赋值
	var loginResMes message.LoginResMes
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	//合法
	//	loginResMes.Code = 200
	//} else {
	//	//不合法
	//	//fmt.Println("用户名 密码 不合法")
	//	loginResMes.Code = 500
	//	loginResMes.Error = "请注册再使用"
	//
	//}

	//使用model.MyUserDao.Login()到redis验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		//～用户登录成功，将该登录成功用户放入到UserMgr中
		// 但是缺少userId
		this.UserId = loginMes.UserId
		// 作为一个功能点，独立存在于userMgr.go中
		userMgr.AddOnlineUsers(this)
		// 通知其他的在线用户
		this.NotifyOthersOnlineUser(loginMes.UserId)
		// 将当前在线用户返回给登录用户
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id) //UsersId切片
		}
		fmt.Println(user, "登录成功")
	}
	//3.序列化LoginResMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	//4.将data赋值给resMes的Data
	resMes.Data = string(data) // string resMes.Data []byte data
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	// 5.发送数据
	// 因为修改为分层模式，需要创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessRegister(mes message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}

	// 处理注册，并且发送给客户端
	// 1.声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//2.声明一个registerResMes，并完成赋值
	var registerResMes message.LoginResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()

		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
			fmt.Println(err.Error())
		}
	} else {
		registerResMes.Code = 200
	}

	// 将registerResMes进行反序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	//4.将data赋值给resMes的Data
	resMes.Data = string(data) // string resMes.Data []byte data
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	// 5.发送数据
	// 因为修改为分层模式，需要创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
