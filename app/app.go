package app

import (
	"fornever.org/app/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// CreateApp but not run
func CreateApp(param *WebAppParam) *WebApplication {
	engine := gin.Default()
	app := &WebApplication{
		param:  param,
		engine: engine,
	}
	var db *gorm.DB

	if len(param.SqliteDsn) > 0 {
		db, _ = gorm.Open(sqlite.Open(param.SqliteDsn), &gorm.Config{})
	}
	if len(param.MysqlDsn) > 0 {
		db, _ = gorm.Open(mysql.Open(param.MysqlDsn), &gorm.Config{})
	}
	if len(param.PgDsn) > 0 {
		db, _ = gorm.Open(postgres.Open(param.PgDsn), &gorm.Config{})
	}

	_ = db.AutoMigrate(&model.Schedule{}, &model.Task{})
	app.mount()
	return app
}

type WebApplication struct {
	param  *WebAppParam
	engine *gin.Engine
}

func (app *WebApplication) mount() {
	app.engine.GET("/health", app.health)

}

func (app *WebApplication) Run(addr string) error {
	return app.engine.Run(addr)
}
