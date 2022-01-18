package apiserver

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

// TestTable ...
type table struct {
	id     int
	status string
}

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info("Server is starting")

	return http.ListenAndServe(s.config.BindAddr, s.router)

}

// ConfigureLogger ...
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

// ConfigureRouter ...
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.HandleHello())
	s.router.HandleFunc("/DBTest", s.DBTest())

}

// HandleHello ...
func (s *APIServer) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

// DBTest ...
func (s *APIServer) DBTest() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		db, err := sql.Open("postgres", "host=localhost dbname=restapi_dev sslmode=disable")
		if err != nil {
			s.logger.Fatal(err)
			os.Exit(1)
		}
		defer func() { _ = db.Close() }()
		defer func() { _ = req.Body.Close() }()
		rows, err := db.Query("SELECT id, status from status")
		if err != nil {
			s.logger.Fatal(err)
			os.Exit(1)
		}
		defer func() { _ = rows.Close() }()
		statuses := []table{}
		for rows.Next() {
			status := table{}
			err := rows.Scan(&status.id, &status.status)
			if err != nil {
				s.logger.Fatal(err)
				os.Exit(1)
			}
			statuses = append(statuses, status)
		}
		for _, status := range statuses {
			io.WriteString(res, string(status.id)+status.status)
		}
	}
}
