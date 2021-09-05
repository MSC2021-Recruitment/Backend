package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	CreatorRefer uint
	Creator User `gorm:"foreignKey:CreatorRefer"`

	Title string `gorm:"unique"`
	Content string
	Group string
}
