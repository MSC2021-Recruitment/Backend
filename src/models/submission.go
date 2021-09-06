package models

import (
	"time"

	"gorm.io/gorm"
)

type Submission struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"created-at"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	UserID     uint           `json:"user-id"`
	QuestionID uint           `json:"question-id"`
	Content    string         `json:"content"`
}
