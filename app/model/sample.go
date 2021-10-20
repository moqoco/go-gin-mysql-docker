package model

import (
	"gorm.io/gorm"
)

type Sample struct {
	ID   uint	`json:"id"`
	Txt  string `json:"txt"`
}

func (s *Sample) First() (g *gorm.DB) {
	return DB.First(&s)
}