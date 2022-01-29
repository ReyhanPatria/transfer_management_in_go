package entity

import (
	"time"
)

type Product struct {
	Key         string    `gorm:"column:KEY" json:"key"`
	Type        string    `gorm:"column:TYPE" json:"type"`
	MinLimit    int       `gorm:"column:MINLIMIT" json:"minLimit"`
	MaxLimit    int       `gorm:"column:MAXLIMIT" json:"maxLimit"`
	Fee         int       `gorm:"column:FEE" json:"fee"`
	CutOn       time.Time `gorm:"column:CUTON" json:"cutOn"`
	CutOff      time.Time `gorm:"column:CUTOFF" json:"cutOff"`
	Description string    `gorm:"column:DESCRIPTION" json:"description"`
	KeepDay     int       `gorm:"column:KEEPDAY" json:"keepDay"`
}
