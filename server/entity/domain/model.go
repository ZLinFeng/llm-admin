package domain

import "time"

type GlobalDBModel struct {
	ID       uint      `gorm:"primarykey" json:"ID"` //主键
	CreateAt time.Time //创建时间
	UPDATEAt time.Time // 更新时间
}
