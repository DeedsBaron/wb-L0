package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
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
	s.router.HandleFunc("/{id}", s.RenderTemplate()).Methods("GET")
	//s.router.PathPrefix("/").Handler(http.FileServer(http.Dir("./interface/")))
	//s.router.PathPrefix("/{id}").Handler(http.StripPrefix("/{id}", http.FileServer(http.Dir("./interface"))))

}

func (s *APIServer) RenderTemplate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("here")
		tmpl, _ := template.ParseFiles("interface/index.html")
		tmpl.Execute(w, 1)
		//fmt.Fprintf(w, "<b>Main Text</b>")
		//err := s.storage.GetDetailedStats(context.Background(), w, mux.Vars(r)["id"])
		//if err != nil {
		//	utils.HttpErrorWithoutBackSlashN(w, err.Error(), http.StatusBadRequest)
		//	s.logger.Error(err.Error())
		//	return
		//}
		//s.logger.Debug("GET GetDetailedStats method SUCCESS")
	}
}
