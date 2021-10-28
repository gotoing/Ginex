package global

import (
	"ginex/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Setting
	//GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger
	GVA_DB  *gorm.DB
)
