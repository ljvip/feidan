package global

import (
	"time"
)

type GVA_MODEL struct {
	ID         uint `gorm:"primarykey" json:"ID"` // 主键ID
	CreateTime time.Time
	UpdateTime time.Time // 更新时间
}
