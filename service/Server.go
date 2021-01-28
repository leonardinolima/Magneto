package service

import (
	"github.com/gorilla/mux"
	"github.com/magneto/routes"
	"log"
	"net/http"
	"time"
)

type Server struct {

}

func (p *Server) RunServer()  {
	routes := new(routes.Routes)
	port:=":8080"
	router := mux.NewRouter()
	routes.SetupRoutes(router)
	server := &http.Server{
		Addr:              port,
		Handler:           router,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
	server.ListenAndServe()
}
