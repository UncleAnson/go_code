package process2

import "fmt"

var userMgr *UserMgr

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 利用初始化函数对userMgr初始化
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUsers的添加
func (this *UserMgr) AddOnlineUsers(process *UserProcess) {
	this.onlineUsers[process.UserId] = process
}

// 完成对onlineUsers的删除
func (this *UserMgr) DelOnlineUsers(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线用户
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

// 根据id查找当前用户
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//从mmap中取出一个值，带检测方式
	up, ok := this.onlineUsers[userId]
	if !ok { // 说明查找的用户不存在（不在线）
		err = fmt.Errorf("用户%d 不在线", userId)
	}
	return
}
