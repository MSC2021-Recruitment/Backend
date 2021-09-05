package models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	SubmitterRefer uint
	Submitter      User `gorm:"foreignKey:SubmitterRefer"`
	QuestionRefer  uint
	Question       Question `gorm:"foreignKey:QuestionRefer"`
	Content        string
}
