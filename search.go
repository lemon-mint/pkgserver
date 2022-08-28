package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/lemon-mint/pkgserver/db"
)

var nullArrJSON = []byte("[]")

func SearchPackagesHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query().Get("q")
	lim := r.URL.Query().Get("limit")
	off := r.URL.Query().Get("offset")
	if lim == "" {
		lim = "50"
	}
	if off == "" {
		off = "0"
	}
	var limit int32
	var offset int32
	if l, err := strconv.Atoi(lim); err != nil || l > 100 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else {
		limit = int32(l)
	}
	if o, err := strconv.Atoi(off); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else {
		offset = int32(o)
	}

	pkgs, err := DBQueries.SearchPackages(context.Background(), db.SearchPackagesParams{
		ToTsquery: query,
		Offset:    offset,
		Limit:     limit,
	})
	if err != nil && err != pgx.ErrNoRows {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if len(pkgs) <= 0 {
		w.WriteHeader(http.StatusOK)
		w.Write(nullArrJSON)
		return
	}

	err = json.NewEncoder(w).Encode(pkgs)
	if err != nil {
		return
	}
}
