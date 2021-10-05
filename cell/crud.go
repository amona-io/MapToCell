package cell

import (
	"gorm.io/gorm"
	"time"
)

type Tabler interface {
	TableName() string
}

type Repository interface {
	Create(DB *gorm.DB)	error
	Delete(DB *gorm.DB) error
}

func (DBCell) TableName() string {
	return "cells"
}

type BaseModel struct {
	ID        uint 			`gorm:"primarykey" gorm:"autoIncrement" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DBCell struct {
	BaseModel	BaseModel	`gorm:"embedded"`
	LeftTop			string
	RightTop		string
	LeftBottom		string
	RightBottom		string
	CenterLatitude 	float64
	CenterLongitude float64
	IsInRange		bool
	CenterCity		string
}

func (c *Cell) Create(DB *gorm.DB) error {
	err := DB.Create(c).Error
	return err
}

func (c *DBCell) Delete(DB *gorm.DB) error {
	err := DB.Delete(c).Error
	return err
}

func GetCellsByRange(DB *gorm.DB, centerLatitude, centerLongitude float64, Range float64) []DBCell {
	minLatitude := centerLatitude-Range
	maxLatitude := centerLatitude+Range
	minLongitude := centerLongitude-Range
	maxLongitude := centerLongitude+Range
	result := []DBCell{}
	DB.Where("center_latitude BETWEEN ? AND ?", minLatitude, maxLatitude).
		Where("center_longitude BETWEEN ? AND ?", minLongitude, maxLongitude).Find(&result)
	return result
}