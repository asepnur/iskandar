package main

import (
	"flag"
	"log"

	"github.com/asepnur/iskandar/src/util/conn"
	"github.com/asepnur/iskandar/src/util/env"
	"github.com/asepnur/iskandar/src/util/jsonconfig"
	"github.com/asepnur/iskandar/src/webserver"
)

type configuration struct {
	Database  conn.DatabaseConfig `json:"database"`
	Redis     conn.RedisConfig    `json:"redis"`
	Webserver webserver.Config    `json:"webserver"`
}

func main() {
	flag.Parse()

	// load config
	cfgenv := env.Get()
	config := &configuration{}
	isLoaded := jsonconfig.Load(&config, "/etc/iskandar", cfgenv) || jsonconfig.Load(&config, "./files/etc/iskandar", cfgenv)
	if !isLoaded {
		log.Fatal("Failed to load configuration")
	}
	// initialize instance
	conn.InitRedis(config.Redis)
	conn.Consume("180204", "iskandar")
	conn.InitDB(config.Database)
	webserver.Start(config.Webserver)
}
