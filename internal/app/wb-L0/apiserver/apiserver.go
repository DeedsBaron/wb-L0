package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"wb-L0/internal/app/wb-L0/config"
	"wb-L0/internal/app/wb-L0/logger"
	"wb-L0/internal/app/wb-L0/storage"
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
	s.configureRouter()
	logger.Log.Info("Starting API server")
	return http.ListenAndServe(config.Config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/{id}", s.RenderTemplate()).Methods("GET")
	s.router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./interface/css"))))
}

func (s *APIServer) Test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Here")
	}
}

func (s *APIServer) RenderTemplate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storage.Cash.Mu.Lock()
		if _, ok := storage.Cash.Store[mux.Vars(r)["id"]]; ok {
			tmpl, err := template.ParseFiles("interface/index.html")
			if err != nil {
				storage.Cash.Mu.Unlock()
				http.Error(w, "Cant parse template!", http.StatusInternalServerError)
				return
			}
			storage.Cash.Mu.Unlock()
			tmpl.Execute(w, storage.Cash.Store[mux.Vars(r)["id"]])
			return
		}
		tmpl, err := template.ParseFiles("interface/not_found.html")
		if err != nil {
			storage.Cash.Mu.Unlock()
			fmt.Println(err.Error())
			http.Error(w, "Cant parse template!", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		err = tmpl.Execute(w, mux.Vars(r)["id"])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		storage.Cash.Mu.Unlock()
	}
}
