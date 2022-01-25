package cell

import (
	"gorm.io/gorm"
	"time"
)

type Tabler interface {
	TableName() string
}

type Repository interface {
	Create(DB *gorm.DB) error
	Delete(DB *gorm.DB) error
}

func (DBCell) TableName() string {
	return "cells"
}

type BaseModel struct {
	ID        uint `gorm:"primarykey" gorm:"autoIncrement" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DBCell struct {
	BaseModel   BaseModel `gorm:"embedded"`
	LeftTop     string
	RightTop    string
	LeftBottom  string
	RightBottom string
	CenterX     float64
	CenterY     float64
	IsInRange   bool
	CenterCity  string
	Deactivate  bool
}

func (c *Cell) Create(DB *gorm.DB) error {
	err := DB.Create(c).Error
	return err
}

func (c *DBCell) Delete(DB *gorm.DB) error {
	err := DB.Delete(c).Error
	return err
}

func GetCellsByRange(DB *gorm.DB, centerX, centerY float64, Range float64) []DBCell {
	minX := centerX - Range
	maxX := centerX + Range
	minY := centerY - Range
	maxY := centerY + Range
	result := []DBCell{}
	DB.Where("center_x BETWEEN ? AND ?", minX, maxX).
		Where("center_y BETWEEN ? AND ?", minY, maxY).Find(&result)
	return result
}
