package models

import "time"

type File struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `json:"userId"`
	OriginalName string    `json:"originalName"`
	StoredName   string    `json:"storedName"`
	Extension    string    `json:"extension"`
	MimeType     string    `json:"mimeType"`
	Size         int64     `json:"size"`
	URL          string    `json:"url"`
	CreatedAt    time.Time `json:"createdAt"`
}
