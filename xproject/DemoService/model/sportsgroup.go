package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	GROUP_MEMBER_DB_NAME          = "group_sports"
	GROUP_MEMBER_GROUP_TABLE_NAME = "group_sports_group"
)

type SportsGroup struct {
	Id     int64
	UserId string
	Icon   string
	Name   string
}

func (sg SportsGroup) TableName() string {
	return fmt.Sprintf("%s.%s", GROUP_MEMBER_DB_NAME, GROUP_MEMBER_GROUP_TABLE_NAME)
}
