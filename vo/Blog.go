package vo

import (
	"time"
)

type BlogType struct {
	Id         int       //ID
	UserId     int       //用户ID
	Title      string    //分类标题
	status     int       //状态：0 初始，1 有效，2无效
	CreateTime time.Time //创建时间
	ModifyTime time.Time //修改时间
}

type Blog struct {
	Id         int       //ID
	UserId     int       //用户ID
	BlogTypeId int       //分类ID
	Title      string    //标题
	Content    string    //内容
	status     int       //状态：0 初始，1 有效，2无效
	CreateTime time.Time //创建时间
	ModifyTime time.Time //修改时间
}

type BlogComment struct {
	Id         int       //ID
	BlogId     int       //博客ID
	UserId     int       //评论人ID
	Content    string    //评论内容
	status     int       //状态：0 初始，1 有效，2无效
	CreateTime time.Time //创建时间
	ModifyTime time.Time //修改时间
}
