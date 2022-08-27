package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
)

func SearchPackagesHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query().Get("q")
	pkgs, err := DBQueries.SearchPackages(context.Background(), query)
	if err != nil && err != pgx.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(pkgs)
	if err != nil {
		return
	}
}
