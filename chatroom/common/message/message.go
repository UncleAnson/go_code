package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string // 消息类型
	Data string
}

type LoginMes struct {
	UserId   int
	UserPwd  string
	UserName string
}

type LoginResMes struct {
	Code     int //500表示用户未注册，200表示登录成功
	Error    string
	//UserName string
}

type RegisterMes struct {
	//...
}
