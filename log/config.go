package log

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Log_level string   `yaml:"log_level"`
}

func init()  {
	viper.SetDefault("log_level", "info")

	flag.String( "log_level", "info", "Log level for metad print out: debug|info|warning")
	viper.BindPFlag("log_level", flag.Lookup("log_level"))

}