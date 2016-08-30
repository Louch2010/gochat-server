package rpc

import (
	"../dao"
	"../vo"
)

/**
* 用户服务
 */
type UserService struct{}

func (this UserService) IsUserExist(user vo.User, resp *Response) error {

}
func (this UserService) AddUser(user vo.User, resp *Response) error {

}
func (this UserService) UserLogin(user vo.User, resp *Response) error {

}
func (this UserService) GetUser(user vo.User, resp *vo.User) error {

}
func (this UserService) GetUserList(user vo.User, resp *[]vo.User) error {

}

/**
* 好友服务
 */
type FriendService struct{}

func (this FriendService) IsFriendExist(friend vo.Friend, resp *Response) error {

}
func (this FriendService) AddFriend(friend vo.Friend, resp *Response) error {

}
func (this FriendService) GetFriend(friend vo.Friend, resp *vo.Friend) error {

}
func (this FriendService) GetFriendList(friend vo.Friend, resp *[]vo.Friend) error {

}
