package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

// Info from config file
type Config struct {
	LogDir     string
	IP         string
	Port       string
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
	UploadDir  string
	CookieID   string
	FormDBUsername string
        FormDBPassword string
        FormDBName     string

}

// Reads info from config file
func ReadConfig() Config {
	configFile := flag.String("config", "/etc/formedit/formedit.conf", "Path to config file")
	flag.Parse()
	var configfile = *configFile
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}
	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
