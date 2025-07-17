package entities

import (
	"time"
)

// Stadium represents a soccer stadium entity
type Stadium struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName specifies the table name for Stadium
func (Stadium) TableName() string {
	return "stadium"
} 