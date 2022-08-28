package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HealthZHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}
