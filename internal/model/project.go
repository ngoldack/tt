package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Active   bool
	Name     string `gorm:"unique"`
	TicketNr string
	Comment  string
}
