package executiongenerator

import (
	"fornever.org/app/model"
	"gorm.io/gorm"
)

// Executions Generator for scheduling
func GenerateExecution(schedule *model.Schedule) (executions []*model.Execution) {

	return
}

func PollExecutions(db *gorm.DB) {
	var schedules []model.Schedule
	db.Find(&schedules)
	for {
		// TODO: refresh schedules ondemand and block too frequent generation
		for _, schedule := range schedules {
			executions := GenerateExecution(&schedule)
			db.Transaction(func(tx *gorm.DB) error {
				db.Create(&executions)
				return nil
			})
		}
	}

}
