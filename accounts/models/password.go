package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Password struct {
	gorm.Model

	Email                  string    `gorm:"type:varchar(100);unique_index;json:email"`
	Mobile                 string    `json:"mobile"`
	Password               string    `json:"password"`
	CurrentSessionAttempts int       `json:"current_session_attempts"`
	LastSessionAttempt     time.Time `json:"email"`
	OverallAttempts        int       `json:"overall_attempts"`
}
