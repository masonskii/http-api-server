package main

import (
	"flag"
	"log"
	"main/internal/apiserver"
	//"encoding/json"
	"github.com/BurntSushi/toml"
)

const (
	reqSelect = "select * from FabProjects.status"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

//var (
//listenAddress = ":8080"
//postAddress = "localhost:8081"
//)

/*
func main() {

	logger init
	logger := setupLogger()
	level.Debug(logger).Log("initLog")

	// handle functions


	// Ot dolbaeba -- dlya dolbaeba
	level.Debug(logger).Log("listener start")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		level.Debug(logger).Log("error, exit")
		os.Exit(1)
	}
}
func setupLogger() (logger log.Logger) {
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = level.NewFilter(logger, level.AllowDebug())
	return
}

*/
func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
