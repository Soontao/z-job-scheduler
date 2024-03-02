package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Schedule struct {
	ID         string `gorm:"primaryKey;type:varchar;size:36;not null"` // UUID
	Name       string `gorm:"type:varchar;size:255;not null"`           // Task Name
	Expression string `gorm:"type:varchar;size:255;not null"`           // Cron Expression

	Created int64 `gorm:"autoCreateTime"`
	Updated int64 `gorm:"autoUpdateTime"`
}

func (s *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}
