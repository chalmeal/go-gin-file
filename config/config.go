package config

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
)

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		fmt.Println("loade config err")
	}
}

// error level normal
func GetInitKey(sec, key string) string {
	k, err := Cfg.Section(sec).GetKey(key)
	if err != nil {
		log.Println(err)
		return ""
	}
	return k.String()
}

// error level fatal
func GetInitKeyFatal(sec, key string) string {
	k, err := Cfg.Section(sec).GetKey(key)
	if err != nil {
		log.Fatal(err)
	}
	return k.String()
}
