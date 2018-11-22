package main

import (
	"github.com/petterhg/go-webapp/daemon"
)
var assetsPath string

func processFlags() *daemon.Config {
	config := &daemon.Config{}
	
	flag.StringVar(&config.ListenSpec)
}