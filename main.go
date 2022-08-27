package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/lemon-mint/envaddr"
	"github.com/lemon-mint/godotenv"
	"v8.run/go/exp/signal2"
	"v8.run/go/exp/util/env"
)

func main() {
	godotenv.Load()

	err := connectPool(env.GetEnvOrDefault("DATABASE_URL", ""))
	if err != nil {
		panic(err)
	}

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
	router.GET("/__api/v1/search", SearchPackagesHandler)

	srv := http.Server{Handler: router, IdleTimeout: time.Second * 25}
	go srv.Serve(ln)
	signal2.WFI()

	// Graceful shutdown
	err = srv.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
}
