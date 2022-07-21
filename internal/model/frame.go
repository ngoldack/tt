package model

import (
	"gorm.io/gorm"
	"time"
)

type Frame struct {
	gorm.Model
	Active    bool
	StartTime time.Time
	StopTime  time.Time
	projectID int
	Project   Project
	tagId     int
	Tag       Tag
	Comment   string
}
