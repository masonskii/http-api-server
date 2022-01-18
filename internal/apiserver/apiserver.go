package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

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

}

// HandleHello ...
func (s *APIServer) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
