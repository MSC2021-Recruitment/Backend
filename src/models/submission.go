package models

import (
	"gorm.io/gorm"
	"time"
)

type Submission struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time      `json:"created-at"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	SubmitterRefer uint           `json:"-"`
	Submitter      User           `gorm:"foreignKey:SubmitterRefer"`
	QuestionRefer  uint           `json:"-"`
	Question       Question       `gorm:"foreignKey:QuestionRefer"`
	Content        string         `json:"content"`
}
