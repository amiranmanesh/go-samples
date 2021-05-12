package models

import "time"

type ModelID struct {
	ID uint `gorm:"primary_key" json:"id"`
}

type ModelTimeStamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ModelDeletedAt struct {
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
