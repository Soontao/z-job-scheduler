package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Execution struct {
	ID string `gorm:"primaryKey;type:varchar;size:36;not null"` // UUID

	Status int    `gorm:"type:int;not null"`             // 0: pending, 1: running, 2: success, 3: failed, 4: canceled
	TaskID string `gorm:"type:varchar;size:36;not null"` // referenced Task ID

	Created int64 `gorm:"autoCreateTime:milli"`
	Updated int64 `gorm:"autoUpdateTime:milli"`
}

type ExecutionLog struct {
	ID          string `gorm:"primaryKey;type:varchar;size:36;not null"`
	ExecutionID string `gorm:"type:varchar;size:36;not null"` // referenced Execution
	Level       int    `gorm:"type:int;not null"`             // 0: debug, 1: info, 2: warn, 3: error
	Msg         string `gorm:"type:varchar;size:1000"`
}

func (e *Execution) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	return
}

func (l *ExecutionLog) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.NewString()
	return
}
