package account

import "feidan-api/global"

type PlatformLog struct {
	global.GVA_MODEL
	AdminUserId int32  `gorm:"column:admin_user_id"`
	Url         string `gorm:"column:url"`
	Username    string `gorm:"column:username"`
	Send        []byte `gorm:"column:send"`
	Input       []byte `gorm:"column:input"`
}

func (s PlatformLog) TableName() string {
	return "platform_log"
}
