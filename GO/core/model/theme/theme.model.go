package thememodel

import "time"

// Theme stores the theme values for boxes
type Theme struct {
	ID        int
	Name      string
	Value     string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt time.Time
}
