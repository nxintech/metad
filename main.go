package main

import (
	flag "github.com/spf13/pflag"
	"fmt"
	"github.com/yunify/metad/log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
	"github.com/spf13/viper"
	"io/ioutil"
	"strconv"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Error("Main Recover: %v, try restart.", r)
			time.Sleep(time.Duration(1000) * time.Millisecond)
			main()
		}
	}()

	flag.Parse()

	if viper.GetBool("version") {
		fmt.Printf("Metad Version: %s\n", VERSION)
		fmt.Printf("Git Version: %s\n", GIT_VERSION)
		fmt.Printf("Go Version: %s\n", runtime.Version())
		fmt.Printf("Go OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	if viper.GetBool("pprof") {
		fmt.Printf("Start pprof, 127.0.0.1:6060\n")
		go func() {
			log.Fatal("%v", http.ListenAndServe("127.0.0.1:6060", nil))
		}()
	}

	var config *Config
	var err error
	if config, err = NewConfig(); err != nil {
		log.Fatal(err.Error())
		os.Exit(-1)
	}

	// Update config from commandline flags.

	if config.Log.Log_level != "" {
		println("set log level to:", config.Log.Log_level)
		log.SetLevel(config.Log.Log_level)
	}

	if config.Pid_file != "" {
		log.Info("Writing pid %d to %s", os.Getpid(), config.Pid_file)
		if err := ioutil.WriteFile(config.Pid_file, []byte(strconv.Itoa(os.Getpid())), 0644); err != nil {
			log.Fatal("Failed to write pid file %s: %v", config.Pid_file, err)
		}
	}



	log.Info("Starting metad %s", VERSION)
	metad, err = NewMetad(config)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(-1)
	}

	metad.Init()
	metad.Serve()
}
