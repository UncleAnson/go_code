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
	Conn net.Conn
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
