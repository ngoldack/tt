package model

import (
	"gorm.io/gorm"
	"time"
)

type Frame struct {
	gorm.Model
	Active    bool
	StartTime time.Time `gorm:"unique"`
	StopTime  time.Time
	ProjectID int
	Project   *Project
	TagId     int
	Tag       *Tag
	Comment   string
}
