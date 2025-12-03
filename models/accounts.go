package models

import "time"

type Account struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"` // we’ll change this to hashed later
	CreatedAt time.Time
	UpdatedAt time.Time
}
