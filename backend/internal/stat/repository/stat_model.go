package repository

import (
	"time"

	"gorm.io/gorm"
)

type ClickModel struct {
	gorm.Model
	LinkID   uint     `gorm:"index"`
	UA       UAModel  `gorm:"foreignKey:ClickID;constraint:OnDelete:CASCADE"`
	UTM      UTMModel `gorm:"foreignKey:ClickID;constraint:OnDelete:CASCADE"`
	GEO      GEOModel `gorm:"foreignKey:ClickID;constraint:OnDelete:CASCADE"`
	Datetime time.Time
}

func (ClickModel) TableName() string { return "clicks" }

type UTMModel struct {
	ID       uint   `gorm:"primarykey"`
	ClickID  uint   `gorm:"index"`
	Term     string `gorm:"default:null"`
	Source   string
	Medium   string
	Content  string `gorm:"default:null"`
	Campaign string
}

func (UTMModel) TableName() string { return "utm" }

type GEOModel struct {
	ID      uint `gorm:"primarykey"`
	ClickID uint `gorm:"index"`
	City    string
	Region  string
	Country string
}

func (GEOModel) TableName() string { return "geo" }

type UAModel struct {
	ID      uint `gorm:"primarykey"`
	ClickID uint `gorm:"index"`
	OS      string
	Device  string
	Browser string
}

func (UAModel) TableName() string { return "ua" }
