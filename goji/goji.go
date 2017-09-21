package goji

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"utils"

	goji "github.com/zenazn/goji/web"
)

type Server struct {
	ws   *goji.Mux
	port int
}

func NewServer(port int) *Server {
	server := &Server{
		ws:   goji.New(),
		port: port,
	}
	flag.Set("bind", ":"+fmt.Sprintf("%v", port))
	return server
}

func wrap(handler utils.Handler) func(goji.C, http.ResponseWriter, *http.Request) {
	return func(c goji.C, response http.ResponseWriter, request *http.Request) {
		result, err := handler(request)
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}
		utils.WriteJson(response, result)
	}
}

func (s *Server) AttachRoute(route string, handler utils.Handler) {
	s.ws.Get(route, wrap(handler))
}

func (s *Server) Run() error {
	return http.ListenAndServe(":"+strconv.Itoa(s.port), s.ws)
}
