package main

import (
	// "errors"

	"plugin"

	"strings"

	"encoding/gob"
	"os"

	"github.com/gin-gonic/gin"
	pak "github.com/septianw/shiny-telegram/experiment01/sharedpak"
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

// This function will load *.so library without parsing its function.
// After load library with this function you need to lookup your function.
func LoadSo(path string) *plugin.Plugin {
	plug, err := plugin.Open(path)
	pak.ErrHandler(err)

	return plug
}

func LoadCoreModule(moduleName string) *Module {
	var mod Module
	modpath := strings.Join([]string{Modloc, "core", moduleName, moduleName + ".so"}, "/")

	lib := LoadSo(modpath)
	bootsym, err := lib.Lookup("Bootstrap")
	pak.ErrHandler(err)

	routersym, err := lib.Lookup("Routers")
	pak.ErrHandler(err)

	mod.Bootstrap = bootsym.(func())
	mod.Router = routersym.(func(*gin.Engine))

	return &mod
}

func LoadContribModule(moduleName string) *Module {
	var mod Module
	modpath := strings.Join([]string{Modloc, "contrib", moduleName, moduleName + ".so"}, "/")

	lib := LoadSo(modpath)
	bootsym, err := lib.Lookup("Bootstrap")
	pak.ErrHandler(err)

	routersym, err := lib.Lookup("Routers")
	pak.ErrHandler(err)

	mod.Bootstrap = bootsym.(func())
	mod.Router = routersym.(func(*gin.Engine))

	return &mod
}

func WriteRuntime(rt pak.Runtime) {
	RuntimeFile, err := os.OpenFile("/tmp/shinyRuntimeFile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	pak.ErrHandler(err)

	enc := gob.NewEncoder(RuntimeFile)
	err = enc.Encode(rt)
	pak.ErrHandler(err)

	err = RuntimeFile.Close()
	pak.ErrHandler(err)
}
