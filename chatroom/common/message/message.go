package message

import "golang_code/src/chatroom/server/model"

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

//～定义用户状态常量
const (
	UserOnline = iota
	UserOffline
	UserBustStatus
)

type Message struct {
	Type string // 消息类型
	Data string
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code    int    `json:"code"` //500表示用户未注册，200表示登录成功
	Error   string `json:"error"`
	UsersId []int  //增加字段，保存用户id的切片
	//UserName string
}

type RegisterMes struct {
	//...升级用User结构体表示
	User model.User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"` //400表示用户已被占用，200表示注册成功
	Error string `json:"error"`
}

// 为了配合服务器端推送用户状态的通知
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
