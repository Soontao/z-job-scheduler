package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Schedule struct {
	ID         string `gorm:"primaryKey;type:varchar;size:36;not null"` // UUID
	Name       string `gorm:"type:varchar;size:255;not null"`           // Task Name
	Expression string `gorm:"type:varchar;size:255;not null"`           // Expression or Type
	Params     string `gorm:"type:varchar;size:1000"`                   // free params

	Created int64 `gorm:"autoCreateTime:milli"`
	Updated int64 `gorm:"autoUpdateTime:milli"`
}

func (s *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}
