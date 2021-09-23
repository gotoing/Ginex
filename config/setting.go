package config

type Setting struct {
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	//JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	//Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	//Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	//Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	//Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`

	// auto
	//AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	//Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	//AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
	//TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencentCOS" yaml:"tencent-cos"`
	//Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	//Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`
}
