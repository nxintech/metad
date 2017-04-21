package backends

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)
type Config struct {
	Backend        string
	Prefix         string
	Group          string
	Basic_auth     bool
	Client_ca_keys string
	Client_cert    string
	Client_key     string
	Nodes          []string
	Password       string
	Username       string
}

func init()  {
	viper.SetDefault("backend","local")
	viper.SetDefault("prefix", "")
	viper.SetDefault("group", "default")

	flag.String( "backend", "local", "The metad backend type")
	viper.BindPFlag("backend",flag.Lookup("backend"))
	
	flag.StringP( "config", "c","", "The configuration file path")
	viper.BindPFlag("config",flag.Lookup("config"))
	
	flag.String( "group", "default", "The metad's group name, same group share same mapping config from backend")
	viper.BindPFlag("group",flag.Lookup("group"))
	
	flag.Bool( "basic_auth", false, "Use Basic Auth to authenticate (only used with -backend=etcd)")
	viper.BindPFlag("basic_auth", flag.Lookup("basic_auth"))
	
	flag.String( "client_ca_keys", "", "The client ca keys")
	viper.BindPFlag("client_ca_keys",flag.Lookup("client_ca_keys"))
	
	flag.String( "client_cert", "", "The client cert")
	viper.BindPFlag("client_cert",flag.Lookup("client_cert"))
	
	flag.String( "client_key", "", "The client key")
	viper.BindPFlag("client_key",flag.Lookup("client_key"))
	
	flag.StringSliceP("nodes","n",[]string{"localhost:2379"}, "List of backend nodes")
	viper.BindPFlag("nodes",flag.Lookup("nodes"))
	
	flag.StringP( "username", "u","", "The username to authenticate as (only used with etcd ")
	viper.BindPFlag("username",flag.Lookup("username"))
	
	flag.StringP( "password","p", "", "The password to authenticate with (only used with etcd ")
	viper.BindPFlag("password",flag.Lookup("password"))
	
}