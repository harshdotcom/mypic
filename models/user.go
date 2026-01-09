package models

import "time"

type User struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	UserName         string    `gorm:"unique;not null" json:"userName"`
	Name             string    `gorm:"not null" json:"name"`
	Email            string    `gorm:"unique;not null" json:"email"`
	UserPassword     string    `gorm:"not null" json:"-"`
	UserLogoURL      string    `json:"userLogoURL"`
	TimeStamp        time.Time `gorm:"autoCreateTime" json:"timeStamp"`
	UpdatedTimeStamp time.Time `gorm:"autoUpdateTime" json:"updatedTimeStamp"`
}
