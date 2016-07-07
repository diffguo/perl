package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USER_DB_NAME    = "db_user"
	USER_TABLE_NAME = "user_basic_info"
)

type UserBind struct {
	UserId                    int64
	phone_number              string
	phone_number_vertify_time string
	wx_openid                 string
	wx_openid_vertify_time    string
	qq_openid                 string
	qq_openid_vertify_time    string
	weibo_openid              string
	weibo_openid_vertify_time string
	email                     string
	email_vertify_time        string
	create_time               string
	update_time               string
}

type UserBase struct {
	UserId         int64
	bind_num       int8
	nick           string
	portrait       string
	sex            string
	education      string
	nation         string
	blood_type     string
	birthday       string
	certificate_no string
	street         string
	province       string
	city           string
	area           string
	create_time    string
	update_time    string
}

type ShareAction struct {
	ShareId     int64
	TopicId     int64
	PreUserId   int64
	NextUserId  int64
	create_time string
}

type ShareCommunicate struct {
	ShareId     int64
	UserId      int64
	words       string
	create_time string
}

func (user *User) TableName() string {
	return fmt.Sprintf("%s.%s", GROUP_MEMBER_DB_NAME, GROUP_MEMBER_GROUP_TABLE_NAME)
}
