package gorilla

import (
	"net/http"
	"strconv"
	"utils"

	"github.com/gorilla/mux"
)

type Server struct {
	ws   *mux.Router
	port int
}

func NewServer(port int) *Server {
	return &Server{
		ws:   mux.NewRouter(),
		port: port,
	}
}

func wrap(handler utils.Handler) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		result, err := handler(request)
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}
		utils.WriteJson(response, result)
	}
}

func (s *Server) AttachRoute(route string, handler utils.Handler) {
	s.ws.HandleFunc(route, wrap(handler)).Methods("GET")
}

func (s *Server) Run() error {
	return http.ListenAndServe(":"+strconv.Itoa(s.port), s.ws)
}
