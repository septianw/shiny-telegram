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

type Dbconf struct {
	Type     string
	Host     string
	Port     uint16
	User     string
	Pass     string
	Database string
}

type Status struct {
	Installed uint8
}

var Spin = spinner.New(spinner.CharSets[24], 100*time.Millisecond)
var ListenAddr, Dsn string

// var Config

// NOTE: Dari setiap module ada semacam hook yang dapat dipanggil pada bootstrap level berapa.

func RunBootLevel0() {
	var files []string

	// fmt.Println()
	Spin.Start()
	Spin.Suffix = "  Check files existence:"
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
			// fmt.Println("wow")
			os.Exit(1)
		}
		// else {
		// 	log.Println(lcheck + " check.")
		// }
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

	// fmt.Printf("\n%+v\n", viper.Get("schema"))

	// fmt.Printf("Loaded configuration file: %s\n", viper.ConfigFileUsed())

	ListenAddr = fmt.Sprintf("%s:%d", viper.GetString("server.bind"), viper.GetInt("server.port"))
	// fmt.Printf("Listening at %s\n", ListenAddr)

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
	// time.Sleep(10 * time.Second)
	Spin.Stop()
}

func RunBootLevel1() {
	var dbconf Dbconf

	RunBootLevel0()
	Spin.Start()
	Spin.Suffix = "  This is booting level 1"
	// fmt.Printf("server: %+v", GetConfig("server"))

	// basic connectivity
	d := GetConfig("database").(map[string]interface{})

	dbconf.Host = d["hostname"].(string)
	dbconf.Type = d["type"].(string)
	// convert dari map viper ke int64 dan convert lagi ke uint16
	// karena int64 terlalu besar untuk menyimpan port yang isinya maksimum hanya 65535
	dbconf.Port = uint16(d["port"].(int64))
	dbconf.User = d["username"].(string)
	dbconf.Pass = d["password"].(string)
	dbconf.Database = d["database"].(string)

	succeed, errPing := PingDb(dbconf)
	if !succeed {
		log.Fatalln(errPing)
		os.Exit(3)
	}

	if !SetupDb(dbconf) {
		// fmt.Println("Database migration success.")
		fmt.Println("Database migration failed.")
		os.Exit(3)
	}
	// else {
	// fmt.Println("Database connection succeed")
	// }

	// basic connectivity
	//   db
	//   cache
	// basic table structure
	//   check schema structure
	//   schema exist
	// time.Sleep(10 * time.Second)
	Spin.Stop()
}

// TODO: tambahkan config manifest pada setiap module
// TODO: load config config itu dan gunakan viper merge config untuk merge.
// TODO: format config pakai map, lalu loop config tiap module pakai range map.

func RunBootLevel2() {
	RunBootLevel1()
	Spin.Start()
	Spin.Suffix = "  This is booting level 2"

	modloc := viper.GetString("moduleLocation")
	fmt.Printf("%+v", modloc)

	// init core
	//   setup
	//   run
	// collecting module
	// setup basic module
	// time.Sleep(10 * time.Second)
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
	// time.Sleep(10 * time.Second)
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
