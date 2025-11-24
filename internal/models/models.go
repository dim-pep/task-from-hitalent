package models

import "time"

type Question struct {
	ID        int       `gorm:"primaryKey;column:id" json:"id"`
	Text      string    `gorm:"type:text;not null;column:text" json:"text"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	Answers   []Answer  `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE;" json:"answers,omitempty"`
}

type Answer struct {
	ID         int       `gorm:"primaryKey;column:id" json:"id"`
	QuestionID int       `gorm:"not null;column:question_id;index" json:"question_id"`
	UserID     string    `gorm:"type:uuid;not null;column:user_id" json:"user_id"`
	Text       string    `gorm:"type:text;not null;column:text" json:"text"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
