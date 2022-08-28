package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/lemon-mint/pkgserver/db"
)

func CheckToken(t string) bool {
	data, ok := box.Base64Open(t)
	if !ok {
		return false
	}
	exp, err := strconv.Atoi(string(data))
	if err != nil {
		return false
	}

	if time.Now().UTC().Unix() > int64(exp) {
		return false
	}

	return true
}

var OKBytes = []byte("OK")

func AdminCreatePackageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if !CheckToken(token) {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	var req db.CreatePackageParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = DBQueries.CreatePackage(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(OKBytes)
}

func AdminDeletePackageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if !CheckToken(token) {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	type DeletePackageParams struct {
		PkgName string `json:"pkg_name"`
	}
	var req DeletePackageParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = DBQueries.DeletePackage(context.Background(), req.PkgName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(OKBytes)
}
