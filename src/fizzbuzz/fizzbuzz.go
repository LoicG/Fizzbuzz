package main

import (
	"emicklei"
	"flag"
	"fmt"
	"goji"
	"gorilla"
	"log"
	"os"
	"strings"
	"utils"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, strings.TrimSpace(`
Simple fizz-buzz REST server [OPTIONS]
routers & frameworks:
	- gorilla (default): https://github.com/gorilla/mux
	- goji: https://github.com/zenazn/goji/
	- emicklei: https://github.com/emicklei/go-restful
`)+"\n")
		flag.PrintDefaults()
	}
	port := flag.Int("port", 8084, "server port")
	router := flag.String("router", "gorilla", "router")
	flag.Parse()

	var server utils.Server
	switch *router {
	case "gorilla":
		server = gorilla.NewServer(*port)
	case "goji":
		server = goji.NewServer(*port)
	case "emicklei":
		server = emicklei.NewServer(*port)
	default:
		log.Fatalf("Invalid router type: %s", *router)
	}

	fmt.Println("Router: ", *router)
	fmt.Println("Port: ", *port)
	server.AttachRoute("/fizzbuzz", FizzBuzz)
	log.Fatal(server.Run())
}
