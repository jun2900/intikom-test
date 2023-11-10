package model

import "gorm.io/gorm"

// Task represents the Task model.
type Task struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Status      string `gorm:"type:varchar(50);default:'pending'" json:"status"`
}
