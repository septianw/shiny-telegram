package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Dbconf struct {
	Type     string
	Host     string
	Port     uint16
	User     string
	Pass     string
	Database string
}

type Runtime struct {
	AppName        string
	Version        string
	BuildId        string
	Stage          string
	ConfigLocation string
	Dbconf         Dbconf
	Modloc         string
	Libloc         string
}

type Status struct {
	Installed uint8
}

type Exception interface{}

type TryCatchBlock struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

func (tc TryCatchBlock) Do() {
	if tc.Finally != nil {
		defer tc.Finally()
	}
	if tc.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tc.Catch(r)
			}
		}()
	}
	tc.Try()
}

func ReadRuntime() Runtime {
	var out Runtime

	RuntimeFile, err := os.OpenFile("/tmp/shinyRuntimeFile", os.O_RDWR|os.O_CREATE, 0400)
	ErrHandler(err)

	dec := gob.NewDecoder(RuntimeFile)
	err = dec.Decode(&out)
	ErrHandler(err)

	err = RuntimeFile.Close()
	ErrHandler(err)

	return out
}

func ErrHandler(err error) {
	if err != nil {
		log.Println(err)
	}
}
