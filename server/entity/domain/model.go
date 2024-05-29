package domain

import "time"

type GlobalDBModel struct {
	ID       uint      `gorm:"not null;unique;primary_key;autoIncrement"`
	CreateAt time.Time //创建时间
	UPDATEAt time.Time // 更新时间
}
