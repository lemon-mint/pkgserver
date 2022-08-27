package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/lemon-mint/envaddr"
	"v8.run/go/exp/signal2"
)

func main() {
	ln, err := net.Listen("tcp", envaddr.Get(":8080"))
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	log.Printf("App Running on http://%s\n", func() string {
		if strings.HasPrefix(ln.Addr().String(), "[::]:") {
			return "localhost" + strings.TrimPrefix(ln.Addr().String(), "[::]")
		}
		return ln.Addr().String()
	}())

	router := httprouter.New()

	srv := http.Server{
		Handler: router,
	}
	go srv.Serve(ln)
	signal2.WFI()

	// Graceful shutdown
	err = srv.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
}
