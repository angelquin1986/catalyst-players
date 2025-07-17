package entities

import (
	"time"
)

// SeasonStatus represents the status of a season
type SeasonStatus int

const (
	SeasonStatusDraft SeasonStatus = iota
	SeasonStatusActive
	SeasonStatusCompleted
	SeasonStatusCancelled
)

// Season represents a soccer season entity
type Season struct {
	ID        uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	LeagueID  uint         `json:"league_id" gorm:"not null"`
	Name      string       `json:"name" gorm:"size:255;not null"`
	StartsAt  time.Time    `json:"starts_at" gorm:"type:timestamp;not null"`
	EndsAt    time.Time    `json:"ends_at" gorm:"type:timestamp;not null"`
	Status    SeasonStatus `json:"status" gorm:"type:int;default:0"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	League  League  `json:"league,omitempty" gorm:"foreignKey:LeagueID"`
	Teams   []Team  `json:"teams" gorm:"many2many:season_team;"`
	Matches []Match `json:"matches,omitempty" gorm:"foreignKey:SeasonID"`
}

// TableName specifies the table name for Season
func (Season) TableName() string {
	return "season"
}
