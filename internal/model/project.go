package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Active   bool
	Name     string
	TicketNr string
	Comment  string
}
