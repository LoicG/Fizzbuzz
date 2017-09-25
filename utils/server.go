package utils

import (
	"net/http"
)

type Handler func(*http.Request) ([]string, error)

type Server interface {
	AttachRoute(string, Handler)
	Run() error
}
