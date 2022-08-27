package main

import (
	"context"
	"net"
	"net/http"

	"github.com/lemon-mint/envaddr"
	"v8.run/go/exp/signal2"
)

func main() {
	ln, err := net.Listen("tcp", envaddr.Get(":8080"))
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	srv := http.Server{}
	go srv.Serve(ln)
	signal2.WFI()

	// Graceful shutdown
	err = srv.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
}
