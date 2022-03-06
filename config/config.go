package config

/*
@Date:2022/2/7
@Author [Ambi](https://github.com/AmbitionLover)

*/

type Server struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// JWT
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
