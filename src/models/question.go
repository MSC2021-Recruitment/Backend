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
	CreatorID uint           `json:"user-id"`

	Title         string       `gorm:"unique" json:"title"`
	Content       string       `json:"content"`
	Group         string       `json:"group"`
	Submissions   []Submission `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AnsweredUsers []*User      `json:"-" gorm:"many2many:users_questions;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
