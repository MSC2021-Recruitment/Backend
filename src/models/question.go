package models

import (
	"gorm.io/gorm"
	"time"
)

type Question struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	CreatorRefer uint `json:"-"`
	Creator      User `gorm:"foreignKey:CreatorRefer"`

	Title   string `gorm:"unique" json:"title"`
	Content string `json:"content"`
	Group   string `json:"group"`
}
