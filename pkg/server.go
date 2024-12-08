package server

import (
	"fmt"
	"log"
	"net/http"

	"ratelimit/pkg/limiter"
	"ratelimit/pkg/middleware"
	utils "ratelimit/utils"

	"github.com/gorilla/mux"
)

type Server struct {
	Router      *mux.Router
	RateLimiter *limiter.RateLimiter
	MaxRequest  int
	BucketToken int
}

func New(m int, b int, rl *limiter.RateLimiter) *Server {
	return &Server{
		MaxRequest:  m,
		BucketToken: b,
		RateLimiter: rl,
		Router:      mux.NewRouter(),
	}
}

func (s *Server) InitRoutes() {
	s.Router.HandleFunc("/", middleware.SetMiddlewareRatelimit(s.Home, float64(s.BucketToken), float64(s.MaxRequest), s.RateLimiter)).Methods("GET")
}

func (s *Server) Run(httpPort string) {
	fmt.Println("listening to port: ", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, s.Router))
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "This is generated HTTP response from server")
}
