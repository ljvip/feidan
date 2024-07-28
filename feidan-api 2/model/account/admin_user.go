package account

import "feidan-api/global"

type AdminUser struct {
	global.GVA_MODEL
	Username string `gorm:"column:username"`
}

func (s AdminUser) TableName() string {
	return "admin_user"
}
