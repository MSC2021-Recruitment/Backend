package models

import (
	"time"

	"gorm.io/gorm"
)

type Submission struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time      `json:"created-at"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	SubmitterRefer uint           `json:"-"`
	Submitter      User           `gorm:"foreignKey:SubmitterRefer" json:"submitter"`
	QuestionRefer  uint           `json:"-"`
	Question       Question       `gorm:"foreignKey:QuestionRefer" json:"question"`
	Content        string         `json:"content"`
}
