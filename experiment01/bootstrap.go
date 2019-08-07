package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/viper"
)

const BOOTSTRAP_LEVEL_0 = 0
const BOOTSTRAP_LEVEL_1 = 1
const BOOTSTRAP_LEVEL_2 = 2
const BOOTSTRAP_LEVEL_3 = 3

var Spin = spinner.New(spinner.CharSets[24], 100*time.Millisecond)
var ListenAddr string

// NOTE: Dari setiap module ada semacam hook yang dapat dipanggil pada bootstrap level berapa.

func RunBootLevel0() {
	var files []string

	// fmt.Println()
	Spin.Start()
	Spin.Suffix = "  Check files existence:\n"
	// time.Sleep(4 * time.Second)
	files = []string{
		fmt.Sprintf("/etc/%s/config", APPNAME),
		fmt.Sprintf("/etc/%s/config.d/", APPNAME),
		fmt.Sprintf("/usr/local/lib/%s", APPNAME),
		fmt.Sprintf("/usr/local/lib/%s/modules", APPNAME),
	}

	for _, lcheck := range files {
		if _, err := os.Stat(lcheck); os.IsNotExist(err) {
			log.Println(lcheck + " not exist")
			fmt.Println("wow")
			os.Exit(30)
		} else {
			log.Println(lcheck + " check.")
		}
	}

	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", APPNAME))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", APPNAME))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("No config file found.")
			// log.Printf("err: %+v\n", err)
			os.Exit(2)
			// Config file not found; ignore error if desired
		} else {
			log.Fatalln(err)
			os.Exit(2)
			// Config file was found but another error was produced
		}
	}

	// fmt.Printf("server: %+v", viper.Get("server"))
	// fmt.Println(viper.ConfigFileUsed())

	ListenAddr = fmt.Sprintf("%s:%d", viper.GetString("server.bind"), viper.GetInt("server.port"))
	fmt.Printf("Listening at %s\n", ListenAddr)

	switch os.Getenv("STAGE") {
	case "production":
		STAGE = "production"
	case "development":
		fallthrough
	case "testing":
		fallthrough
	default:
		STAGE = "development"
	}
	// check integrity (rely on system, we can't check ourself id)
	// check requirement
	//   paths
	//   config
	//   libraries
	Spin.Stop()
}

func RunBootLevel1() {
	RunBootLevel0()
	Spin.Start()
	Spin.Suffix = "  This is booting level 1"
	// basic connectivity
	//   db
	//   cache
	// basic table structure
	//   check schema structure
	//   schema exist
	Spin.Stop()
}

func RunBootLevel2() {
	RunBootLevel1()
	Spin.Start()
	Spin.Suffix = "  This is booting level 2"
	// init core
	//   setup
	//   run
	// collecting module
	// setup basic module
	Spin.Stop()
}

func RunBootLevel3() {
	RunBootLevel2()
	Spin.Start()
	Spin.Suffix = "  This is booting level 3"
	// init contrib
	//   setup
	//   run
	// setup router
	Spin.Stop()
}

func Bootstrap(level int) {
	switch level {
	case BOOTSTRAP_LEVEL_0:
		RunBootLevel0()
		break
	case BOOTSTRAP_LEVEL_1:
		RunBootLevel1()
		break
	case BOOTSTRAP_LEVEL_2:
		RunBootLevel2()
		break
	case BOOTSTRAP_LEVEL_3:
		RunBootLevel3()
		break
	}
}

func BootstrapAll() {
	Bootstrap(BOOTSTRAP_LEVEL_3)
}
