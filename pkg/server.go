package server

import (
	"fmt"
	"log"
	"net/http"

	"ratelimit/pkg/middleware"
	utils "ratelimit/utils"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

type Server struct {
	Router      *mux.Router
	MaxRequest  int
	BucketToken int
}

func New(m int, b int) *Server {
	return &Server{
		MaxRequest:  m,
		BucketToken: b,
		Router:      mux.NewRouter(),
	}
}

func (s *Server) InitRoutes() {
	s.Router.HandleFunc("/", middleware.SetMiddlewareRatelimit(s.Home, rate.Limit(s.MaxRequest), s.BucketToken)).Methods("GET")
}

func (s *Server) Run(httpPort string) {
	fmt.Println("listening to port: ", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, s.Router))
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "This is generated HTTP response from server")
}
