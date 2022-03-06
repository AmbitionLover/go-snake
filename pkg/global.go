package global

import (
	"github.com/AmbitionLover/go-snake/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

/*
@Date:2021-08-28
@Author Ambi 赵帅

*/
var (
	DB                 *gorm.DB
	VP                 *viper.Viper
	LOG                *zap.Logger
	CONFIG             config.Server
	ConcurrencyControl = &singleflight.Group{} // 并发控制
)
