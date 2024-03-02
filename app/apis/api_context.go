package apis

import (
	"fornever.org/app/param"
	"gorm.io/gorm"
)

type APIBootstrapContext struct {
	DB       *gorm.DB
	AppParam *param.WebAppParam
}
