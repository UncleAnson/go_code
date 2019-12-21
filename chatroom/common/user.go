package model

// 定义一个用户的结构体
type User struct {
	UserId   int    `json: "userId"` // 为了序列化和反序列化成功（结构体的tag要与redis数据库中保存的一致）
	UserPwd  string `json: "userPwd"`
	UserName string `json: "userName"`
}
