package vo

import (
	"time"
)

type User struct {
	Id         int       //ID
	Username   string    //用户名
	Password   string    //密码
	Nickname   string    //昵称
	Birthday   time.Time //生日
	Gender     int       //性别：0 男， 1 女
	FaceImg    string    //头像
	Remark     string    //备注签名
	RegTime    time.Time //注册时间
	status     int       //状态：0 初始，1 有效，2 封号
	CreateTime time.Time //创建时间
	ModifyTime time.Time //修改时间
}

type Friend struct {
	Id         int       //ID
	UserId     int       //用户ID
	FriendId   int       //朋友用户ID
	Remark     string    //备注名
	AddTime    time.Time //添加时间
	status     int       //状态：0 初始，1 有效，2无效
	CreateTime time.Time //创建时间
	ModifyTime time.Time //修改时间
}
