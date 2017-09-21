package emicklei

import (
	"net/http"
	"strconv"
	"utils"

	"github.com/emicklei/go-restful"
)

type Server struct {
	ws   *restful.WebService
	port int
}

func NewServer(port int) *Server {
	server := &Server{
		ws:   &restful.WebService{},
		port: port,
	}
	return server
}

func wrap(handler utils.Handler) restful.RouteFunction {
	return func(request *restful.Request, response *restful.Response) {
		result, err := handler(request.Request)
		if err != nil {
			response.WriteError(http.StatusBadRequest, err)
			return
		}
		response.WriteEntity(result)
	}
}

func (s *Server) AttachRoute(route string, handler utils.Handler) {
	s.ws.
		Path(route).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	s.ws.
		Route(s.ws.GET("/").
			To(wrap(handler)).
			Doc("fizz buzz generation"))
	restful.Add(s.ws)
}

func (s *Server) Run() error {
	return http.ListenAndServe(":"+strconv.Itoa(s.port), nil)
}
