package app

import (
	"fornever.org/app/apis"
	"fornever.org/app/model"
	"fornever.org/app/param"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// CreateApp but not run
func CreateApp(param *param.WebAppParam) *WebApplication {
	engine := gin.Default()
	app := &WebApplication{
		param:  param,
		engine: engine,
	}

	if len(param.SqliteDsn) > 0 {
		app.db, _ = gorm.Open(sqlite.Open(param.SqliteDsn), &gorm.Config{})
	}
	if len(param.MysqlDsn) > 0 {
		app.db, _ = gorm.Open(mysql.Open(param.MysqlDsn), &gorm.Config{})
	}
	if len(param.PgDsn) > 0 {
		app.db, _ = gorm.Open(postgres.Open(param.PgDsn), &gorm.Config{})
	}

	_ = app.db.AutoMigrate(
		&model.Schedule{},
		&model.Task{},
		&model.Execution{},
		&model.ExecutionLog{},
	)
	app.mount()
	return app
}

type WebApplication struct {
	param  *param.WebAppParam
	engine *gin.Engine
	db     *gorm.DB
}

func (app *WebApplication) mount() {
	ctx := &apis.APIBootstrapContext{
		DB:       app.db,
		AppParam: app.param,
	}
	apis.HealthAPIs(app.engine.Group("/health"), ctx)
	apis.JobAPIs(app.engine.Group("/jobs"), ctx)
}

func (app *WebApplication) Run(addr string) error {
	return app.engine.Run(addr)
}
