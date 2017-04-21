package main

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yunify/metad/log"
	"github.com/yunify/metad/backends"
)

type Config struct {
	Pid_file      string   `yaml:"pid_file"`
	Xff           bool     `yaml:"xff"`
	Only_self     bool     `yaml:"only_self"`
	Listen        string   `yaml:"listen"`
	Listen_manage string   `yaml:"listen_manage"`
	Backends      backends.Config
	Log           log.Config
}

func init() {

	viper.SetDefault("Listen", ":80")
	viper.SetDefault("Listen_manage", "127.0.0.1:9611")

	viper.SetEnvPrefix("metad")
	viper.AutomaticEnv()
	viper.AddConfigPath("/etc/metad/")
	viper.AddConfigPath(".")
	viper.SetConfigName("metad")
	err :=viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	flag.BoolP("version","v", false, "Show metad version")
	viper.BindPFlag("version",flag.Lookup("version"))
	flag.Bool("pprof", false, "Enable http pprof, port is 6060")
	viper.BindPFlag("pprof",flag.Lookup("pprof"))
	flag.String( "pid_file", "", "PID to write to")
	viper.BindPFlag("pid_file",flag.Lookup("pid_file"))
	flag.Bool( "xff", false, "X-Forwarded-For header support")
	viper.BindPFlag("xff",flag.Lookup("xff"))
	flag.String( "prefix", "", "Backend key path prefix")
	viper.BindPFlag("prefix",flag.Lookup("prefix"))
	flag.Bool( "only_self", false, "Only support self metadata query")
	viper.BindPFlag("only_self",flag.Lookup("only_self"))
	flag.String( "listen", ":80", "Address to listen to (TCP)")
	viper.BindPFlag("listen",flag.Lookup("listen"))
	flag.String( "listen_manage", "127.0.0.1:9611", "Address to listen to for manage requests (TCP)")
	viper.BindPFlag("listen_manage",flag.Lookup("listen_manage"))

}

func NewConfig() (*Config, error) {

	// create config object
	var config Config

	err := viper.Unmarshal(&config)

	return &config, err
}

