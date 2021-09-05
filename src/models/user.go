package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name      string `json:"name"`
	Password  string `json:"-"`
	StudentID string `gorm:"unique" json:"student-id"`
	Admin     bool   `json:"admin"`
	Major     string `json:"major"`
	Telephone string `gorm:"unique" json:"telephone"`
	Email     string `gorm:"unique" json:"email"`
	QQ        string `gorm:"unique" json:"qq"`
	Level     string `json:"level"`
	Wanted    string `json:"wanted"`
	Intro     string `json:"intro"`
}
