package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Schedule struct {
	ID                    string  `gorm:"primaryKey;type:varchar;size:36;not null"` // UUID
	Name                  string  `gorm:"type:varchar;size:255;not null"`           // Task Name
	Expression            *string `gorm:"type:varchar;size:255;not null"`           // Expression or Type
	Params                *string `gorm:"type:varchar;size:1000"`                   // free params
	Status                bool    `gorm:"type:bool;not null"`                       // true: enabled, false: disabled
	ExecutionDelayPenalty bool    `gorm:"type:bool;not null;default:true"`          // execution delay will have impact to schedule, true: enabled, false: disabled
	MaxConcurrency        int     `gorm:"type:int;not null;default:-1"`             // Concurrency, -1 disabled

	Created int64 `gorm:"autoCreateTime:milli"`
	Updated int64 `gorm:"autoUpdateTime:milli"`
}

func (s *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}
