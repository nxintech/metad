package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
	"github.com/yunify/metad/log"
	"github.com/yunify/metad/backends"
	"github.com/spf13/viper"
	"os"
)

func TestConfigFile(t *testing.T) {
	config := Config{
		Log:           log.Config{ Log_level:"debug", },
		Pid_file:      "/var/run/metad.pid",
		Xff:           true,
		Listen:        ":8080",
		Listen_manage: "127.0.0.1:9611",
		Only_self:     true,
		Backends: backends.Config{
			Backend:"etcd",
			Prefix:"/users/uid1",
			Group:        "default",
			Basic_auth:    true,
			Client_ca_keys: "/opt/metad/client_ca_keys",
			Client_cert:   "/opt/metad/client_cert",
			Client_key:    "/opt/metad/client_key",
			Nodes: []string{"192.168.11.1:2379", "192.168.11.2:2379"},
			Username:     "username",
			Password:     "password",
		},

	}

	data, err := yaml.Marshal(config)
	assert.NoError(t, err)
	configFile, fileErr := os.Create("metad.yml")
	defer configFile.Close()

	fmt.Printf("configFile: %v \n", configFile.Name())

	assert.Nil(t, fileErr)
	c, ioErr := configFile.Write(data)
	assert.Nil(t, ioErr)
	assert.Equal(t, len(data), c)
	configFile.Close()

	config2 := Config{}
	loadErr:= viper.ReadInConfig()
	assert.Nil(t, loadErr)
	//os.Remove("metad.yml")
	viper.Unmarshal(&config2)
	assert.Equal(t, config, config2)
}