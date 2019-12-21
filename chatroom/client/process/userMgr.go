package process

import (
	_ "fmt"
	"golang_code/src/chatroom/common/message"
)

// 客户端要维护的map，但是现在还没初始化
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

// 更新map
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	// 优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		// 新用户
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
}
