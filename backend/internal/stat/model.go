package stat

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Stat struct {
	gorm.Model
	Date   datatypes.Date `json:"date"`
	Clicks int            `json:"clicks"`
	LinkID uint           `json:"link_id" gorm:"index"`
}
