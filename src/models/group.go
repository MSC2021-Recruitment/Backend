package models

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `gorm:"unique" json:"name"`
	Description  string         `gorm:"type:text;" json:"description"`
	Interviewees []*User        `gorm:"many2many:groups_users;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"interviewees"`
	Records      []Interview    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
