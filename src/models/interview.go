package models

import (
	"gorm.io/gorm"
	"time"
)

type Interview struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	GroupID   uint           `json:"group-id"`
	UserID    uint           `json:"user-id"`
	Passed    bool           `json:"passed"`
	Record    string         `json:"record"`
}
