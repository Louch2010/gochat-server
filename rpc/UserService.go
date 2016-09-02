package rpc

import (
	"log"

	"github.com/Louch2010/gochat-server/vo"
)

/**
* 用户服务
 */
type UserService struct{}

//判断是用户是否存在
func (this UserService) IsUserExist(user vo.User, resp *vo.Response) error {
	log.Println("查询用户是否存在，查询条件：", user)
	return nil
}

//添加用户
func (this UserService) AddUser(user vo.User, resp *vo.Response) error {
	return nil
}

//用户登录
func (this UserService) UserLogin(user vo.User, resp *vo.Response) error {
	return nil
}

//获取用户信息
func (this UserService) GetUser(user vo.User, resp *vo.User) error {
	return nil
}

//获取用户列表
func (this UserService) GetUserList(user vo.User, resp *[]vo.User) error {
	return nil
}

/**
* 好友服务
 */
type FriendService struct{}

//判断好友是否存在
func (this FriendService) IsFriendExist(friend vo.Friend, resp *vo.Response) error {
	return nil
}

//添加好友
func (this FriendService) AddFriend(friend vo.Friend, resp *vo.Response) error {
	return nil
}

//获取好友信息
func (this FriendService) GetFriend(friend vo.Friend, resp *vo.Friend) error {
	return nil
}

//获取好友列表
func (this FriendService) GetFriendList(friend vo.Friend, resp *[]vo.Friend) error {
	return nil
}
