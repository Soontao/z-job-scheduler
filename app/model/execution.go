package model

type Execution struct {
	ID string `gorm:"primaryKey;type:varchar;size:36;not null"` // UUID

	Status int    `gorm:"type:int;not null"`             // 0: pending, 1: running, 2: success, 3: failed
	TaskID string `gorm:"type:varchar;size:36;not null"` // referneced Task ID

	Created int64 `gorm:"autoCreateTime:milli"`
	Updated int64 `gorm:"autoUpdateTime:milli"`
}

type ExecutionLog struct {
	ID          string `gorm:"primaryKey;type:varchar;size:36;not null"`
	ExecutionID string `gorm:"type:varchar;size:36;not null"` // referneced Execution
	Level       int    `gorm:"type:int;not null"`             // 0: debug, 1: info, 2: warn, 3: error
	Msg         string `gorm:"type:varchar;size:1000"`
}
