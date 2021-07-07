package main

import (
	"flag"

	"github.com/car12o/audio-mastery/api"
	"github.com/car12o/audio-mastery/internal/audiomastery"
)

type config struct {
	host string
	port uint64
}

var cfg config

func init() {
	flag.StringVar(&cfg.host, "host", "0.0.0.0", "Server host")
	flag.Uint64Var(&cfg.port, "port", 3000, "Server port")
	flag.Parse()
}

func main() {
	app := audiomastery.NewApp()

	server, err := api.NewServer(cfg.host, int(cfg.port), app.Log)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
