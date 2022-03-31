package apiserver

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"wb-L0/internal/app/wb-L0/config"
	"wb-L0/internal/app/wb-L0/logger"
)

var (
	Server *APIServer
)

func init() {
	Server = New()
}

type APIServer struct {
	router *mux.Router
}

func New() *APIServer {
	return &APIServer{
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	time.LoadLocation("Local")
	s.configureRouter()
	logger.Log.Info("Starting API server")
	return http.ListenAndServe(config.Config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/DetailedStats/{id}", s.GetInfo()).Methods("GET")
}

func (s *APIServer) GetInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//err := s.storage.GetDetailedStats(context.Background(), w, mux.Vars(r)["id"])
		//if err != nil {
		//	utils.HttpErrorWithoutBackSlashN(w, err.Error(), http.StatusBadRequest)
		//	s.logger.Error(err.Error())
		//	return
		//}
		//s.logger.Debug("GET GetDetailedStats method SUCCESS")
	}
}
