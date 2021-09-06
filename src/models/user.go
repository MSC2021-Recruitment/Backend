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

	Name      string `json:"name"`                     // 姓名
	Password  string `json:"-"`                        // 密码（哈希存储
	StudentID string `gorm:"unique" json:"student-id"` // 学号
	Admin     bool   `json:"admin"`                    // 是否为管理员（管理员可以发布/查看问题回答情况
	Major     string `json:"major"`                    // 专业
	Telephone string `gorm:"unique" json:"telephone"`  // 手机号
	Email     string `gorm:"unique" json:"email"`      // 邮箱
	QQ        string `gorm:"unique" json:"qq"`         // QQ号
	Level     string `json:"level"`                    // 年级
	Wanted    string `json:"wanted"`                   // 意向组别
	Intro     string `json:"intro"`                    // 自我介绍

	AnsweredQuestions []*Question  `json:"-" gorm:"many2many:users_questions;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 回答的问题，many to many
	Submissions       []Submission `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`                           // 所有的提交 （是否违反设计原则？
}
