package global

import (
	"ginex/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Setting
	//GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger
)
