package db

import (
	"time"
)

type BaseModel struct {
	ID        uint 			`gorm:"primarykey" gorm:"autoIncrement" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Cell struct {
	BaseModel	BaseModel	`gorm:"embedded"`
	LeftTop		string
	RightTop	string
	LeftBottom	string
	RightBottom	string
	Center		string
	IsInRange	bool
	CenterCity	string
}