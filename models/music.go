package models

import "time"

type Music struct {
	ID        uint      `gorm:"primaryKey" json:"id" example:"1"`
	Name      string    `json:"name" example:"Bohemian Rhapsody"`
	Artist    string    `json:"artist" example:"Queen"`
	Album     string    `json:"album" example:"A Night at the Opera"`
	Year      int       `json:"year" example:"1975"`
	Writer    string    `json:"writer" example:"Freddie Mercury"`
	Genre     string    `json:"genre" example:"Rock"`
	Length    int       `json:"length" example:"5"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
}
