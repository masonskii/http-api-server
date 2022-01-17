package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"encoding/json"
	//"fmt"
	"net/http"
	"os"
	//"path/filepath"
	//"time"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

const (
	reqSelect = "select * from FabProjects.status"
)

var (
	listenAddress = ":8080"
	//postAddress = "localhost:8081"
)

type table struct {
	id     int
	status string
}

func main() {

	// logger init
	logger := setupLogger()
	level.Debug(logger).Log("initLog")

	// handle functions
	http.HandleFunc("/DBtest", func(res http.ResponseWriter, req *http.Request) {
		db, err := sql.Open("mysql", "root:0608PAV2002@/FabProjects")
		if err != nil {
			level.Debug(logger).Log("error", err.Error(), "db_conn")
			os.Exit(1)
		}
		defer func() { _ = db.Close() }()
		defer func() { _ = req.Body.Close() }()
		rows, err := db.Query(reqSelect)
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

	// start server
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
