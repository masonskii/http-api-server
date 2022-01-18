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

type table struct {
	id     int
	status string
}

/*
func main() {

	logger init
	logger := setupLogger()
	level.Debug(logger).Log("initLog")

	// handle functions
	http.HandleFunc("/DBtest", func(res http.ResponseWriter, req *http.Request) {
		db, err := sql.Open("mysql", "root:root@/FabProjects")
		if err != nil {
			level.Debug(logger).Log("error", err.Error(), "db_conn")
			os.Exit(1)
		}
		defer func() { _ = db.Close() }()
		defer func() { _ = req.Body.Close() }()
		rows, err := db.Query("SELECT id, status from status")
		if err != nil {
			level.Debug(logger).Log("error", err.Error(), "query")
			os.Exit(1)
		}
		defer func() { _ = rows.Close() }()
		statuses := []table{}
		for rows.Next() {
			status := table{}
			err := rows.Scan(&status.id, &status.status)
			if err != nil {
				level.Debug(logger).Log("error", err.Error(), "scan")
				os.Exit(1)
			}
			statuses = append(statuses, status)
		}
		for _, status := range statuses {
			level.Debug(logger).Log("row", status)
		}

		level.Debug(logger).Log("db test start")
	})

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
