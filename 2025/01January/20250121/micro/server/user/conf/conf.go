package conf

import "github.com/spf13/viper"

type Config struct {
	MySQL
}
type MySQL struct {
	Addr     string `json:"addr" yaml:"addr"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	DB       string `json:"db" yaml:"db"`
}

var (
	Conf *Config
)

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.SetConfigFile("D:\\New\\project\\Golang\\2025\\01January\\20250121\\micro\\server\\user\\conf\\db.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
