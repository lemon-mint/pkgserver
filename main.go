package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/lemon-mint/envaddr"
	"github.com/lemon-mint/godotenv"
	"github.com/lemon-mint/vbox"
	"v8.run/go/exp/signal2"
	"v8.run/go/exp/util/env"
)

func randhex(n int) string {
	buf := make([]byte, n)
	n1, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	if n != n1 {
		panic("n != n1")
	}
	return hex.EncodeToString(buf)
}

var box *vbox.BlackBox
var genToken = flag.Bool("gen_token", false, "keygen")
var tokenExp = flag.Int("exp", 86400*30, "exp")

func main() {
	godotenv.Load()
	flag.Parse()
	box = vbox.NewBlackBox([]byte(env.GetEnvOrDefault("SECRET_KEY", randhex(32))))

	if *genToken {
		fmt.Println(box.Base64Seal([]byte(strconv.Itoa(int(time.Now().UTC().Unix()) + *tokenExp))))
		return
	}

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
	router.GET("/__api/v1/Search", SearchPackagesHandler)
	router.POST("/__api/v1/CreatePackage", AdminCreatePackageHandler)
	router.HEAD("/__api/v1/healthz", HealthZHandler)
	router.GET("/__api/v1/healthz", HealthZHandler)
	router.NotFound = &PackagesHandler{}

	srv := http.Server{Handler: router, IdleTimeout: time.Second * 25}
	go srv.Serve(ln)
	signal2.WFI()

	// Graceful shutdown
	err = srv.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
}
