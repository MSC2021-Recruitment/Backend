package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name      string
	Password  string
	Admin     bool
	Major     string
	Telephone string `gorm:"unique"`
	Email     string `gorm:"unique"`
	QQ        string `gorm:"unique"`
	Level     int32
	Wanted    string
	Intro     string
}
