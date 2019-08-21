package main

import (
	// "errors"

	"log"
	"plugin"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Module struct {
	Bootstrap func()
	Router    func(*gin.Engine)
}

func GetConfig(key string) interface{} {
	return viper.Get(key)
}

func GetAllConfig() map[string]interface{} {
	return viper.AllSettings()
}

func ErrHandler(err error) {
	if err != nil {
		log.Println(err)
	}
}

// This function will load *.so library without parsing its function.
// After load library with this function you need to lookup your function.
func LoadSo(path string) *plugin.Plugin {
	plug, err := plugin.Open(path)
	ErrHandler(err)

	return plug
}

func LoadCoreModule(moduleName string) *Module {
	var mod Module
	modpath := strings.Join([]string{Modloc, "core", moduleName, moduleName + ".so"}, "/")

	lib := LoadSo(modpath)
	bootsym, err := lib.Lookup("Bootstrap")
	ErrHandler(err)

	routersym, err := lib.Lookup("Routers")
	ErrHandler(err)

	mod.Bootstrap = bootsym.(func())
	mod.Router = routersym.(func(*gin.Engine))

	return &mod
}

func LoadContribModule(moduleName string) *Module {
	var mod Module
	modpath := strings.Join([]string{Modloc, "contrib", moduleName, moduleName + ".so"}, "/")

	lib := LoadSo(modpath)
	bootsym, err := lib.Lookup("Bootstrap")
	ErrHandler(err)

	routersym, err := lib.Lookup("Routers")
	ErrHandler(err)

	mod.Bootstrap = bootsym.(func())
	mod.Router = routersym.(func(*gin.Engine))

	return &mod
}

func LoadDatabase(SoPath string) Db {
	var db Db
	log.Println(SoPath)
	So := LoadSo(SoPath)
	PingDb, err := So.Lookup("PingDb")
	ErrHandler(err)
	SetupDb, err := So.Lookup("SetupDb")
	Migrate, err := So.Lookup("MigrateFunc")
	OpenDb, err := So.Lookup("OpenDbFunc")

	db = Db{
		PingDb:  PingDb.(PingDbFunc),
		SetupDb: SetupDb.(SetupDbFunc),
		Migrate: Migrate.(MigrateFunc),
		OpenDb:  OpenDb.(OpenDbFunc),
	}
	// database := SoSym.(DatabaseInterface)
	return db
}
