package goji

import (
	"flag"
	"fmt"
	"net/http"
	"utils"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

type Server struct {
}

func NewServer(port int) *Server {
	server := &Server{}
	flag.Set("bind", ":"+fmt.Sprintf("%v", port))
	return server
}

func wrap(handler utils.Handler) func(web.C, http.ResponseWriter, *http.Request) {
	return func(c web.C, response http.ResponseWriter, request *http.Request) {
		result, err := handler(request)
		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}
		utils.WriteJson(response, result)
	}
}

func (s *Server) AttachRoute(route string, handler utils.Handler) {
	goji.Get(route, wrap(handler))
}

func (s *Server) Run() error {
	goji.Serve()
	return nil
}
