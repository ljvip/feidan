package account

import "feidan-api/global"

type Platform struct {
	global.GVA_MODEL
	Url          string  `gorm:"column:url"`
	Token        string  `gorm:"column:token"`
	Username     string  `gorm:"column:username"`
	Password     string  `gorm:"column:password"`
	PlatformName string  `gorm:"column:platform_name"`
	Redouble     int32   `gorm:"column:redouble"`
	Balance      float64 `gorm:"column:balance"`
	Polling      int32   `gorm:"column:polling"`
	AdminUserId  int32   `gorm:"column:admin_user_id"`
}

func (s Platform) TableName() string {
	return "platform"
}
