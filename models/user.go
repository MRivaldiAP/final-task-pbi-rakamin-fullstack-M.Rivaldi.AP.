package models

import (
    "time"
)

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique;not null"`
    Email     string    `gorm:"unique;not null"`
    Password  string    `gorm:"not null"`
    Photos    []Photo   `gorm:"constraint:OnDelete:CASCADE;"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
