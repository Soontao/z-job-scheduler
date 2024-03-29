package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Task Definition
type Task struct {
	ID   string `gorm:"primaryKey;type:varchar;size:36;not null"` // UUID
	Name string `gorm:"type:varchar;size:255;not null"`           // Task Name

	Created int64 `gorm:"autoCreateTime:milli"`
	Updated int64 `gorm:"autoUpdateTime:milli"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return
}
