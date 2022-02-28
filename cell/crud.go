package cell

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// BitBool is an implementation of a bool for the MySQL type BIT(1).
// This type allows you to avoid wasting an entire byte for MySQL's boolean type TINYINT.
type BitBool bool

// Value implements the driver.Valuer interface,
// and turns the BitBool into a bitfield (BIT(1)) for MySQL storage.
func (b BitBool) Value() (driver.Value, error) {
	if b {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

// Scan implements the sql.Scanner interface,
// and turns the bitfield incoming from MySQL into a BitBool
func (b *BitBool) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	*b = v[0] == 1
	return nil
}

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
	IsInRange   BitBool
	CenterCity  string
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

func GetCellsByMinMax(DB *gorm.DB, minX, minY, maxX, maxY float64) []DBCell {
	fmt.Println("find")
	result := []DBCell{}
	DB.Where("center_x BETWEEN ? AND ?", minX, maxX).
		Where("center_y BETWEEN ? AND ?", minY, maxY).Find(&result)
	return result
}
