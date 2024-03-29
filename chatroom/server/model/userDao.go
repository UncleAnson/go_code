package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 将UserDao实例做成全局变量
var (
	MyUserDao *UserDao
)

// 定义UserDao结构体
type UserDao struct {
	pool *redis.Pool
}

// (last) 使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{pool: pool}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user User, err error) { // 是redis.Conn 不是net.Conn!!
	res, err := redis.String(conn.Do("Hget", "users", id))
	if err != nil {
		//判断错误类型
		if err == redis.ErrNil {
			//表示users哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	// 需要把res反序列化成User实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}
	return
}

// 完成登录的校验
// 1.Login 完成对用户的验证
// 2.如果用户id和密码正确，则返回一个User实例
// 3.如果id或密码有错误，则返回对应的错误
func (this *UserDao) Login(userId int, userPwd string) (user User, err error) {
	//1.先从UserDao连接池中取一个连接
	conn := this.pool.Get()
	defer conn.Close()
	//2.调用 getUserById 查验是否正确
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	// 用户是存在的，但是密码还没校验，需要反序列化User
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	} else {

	}
	return
}

func (this *UserDao) Register(user *User) (err error) { //userId int, userPwd string
	//1. 验证用户名是否存在
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == ERROR_USER_NOTEXISTS {
		// 2. 不存在则允许注册，写入redis数据库
		//1.创建一个User实例
		//user := User{
		//	UserId:   userId,
		//	UserPwd:  userPwd,
		//	UserName: "",
		//}
		data, err := json.Marshal(user)
		if err != nil {
			fmt.Println("json.Marshal err =", err)
		} else {
			fmt.Println("序列化后的数据：", data)
		}
		// 入库
		_, err = conn.Do("Hset", "users", user.UserId, string(data))
		if err != nil {
			fmt.Println("保存注册信息出错 err =", err)
		}
		return nil // 这里的逻辑导致需要返回nil
	} else {
		err = ERROR_USER_EXISTS
		return
	}
}
