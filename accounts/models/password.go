package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Password struct {
	gorm.Model

	Email                  string    `gorm:"type:varchar(100);unique_index;json:email"`
	Mobile                 string    `gorm:"type:varchar(100);unique_index;json:mobile"`
	Password               string    `json:"password"`
	CurrentSessionAttempts int       `gorm:"default:0;json:current_session_attempts"`
	LastSessionAttempt     time.Time `json:"last_session_attempt"`
	OverallAttempts        int       `json:"overall_attempts"`
}
