package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/lemon-mint/pkgserver/db"
	"github.com/lemon-mint/pkgserver/views"
)

func IndexPageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var pkgs []db.Package
	var err error

	lim := r.URL.Query().Get("l")
	off := r.URL.Query().Get("o")
	var limit int32
	var offset int32
	if l, err := strconv.Atoi(lim); err != nil || l > 50 || l <= 0 {
		limit = 25
	} else {
		limit = int32(l)
	}
	if o, err := strconv.Atoi(off); err != nil || o < 0 {
		offset = 0
	} else {
		offset = int32(o)
	}

	query := r.URL.Query().Get("q")
	if query != "" {
		pkgs, err = DBQueries.SearchPackages(context.Background(), db.SearchPackagesParams{
			ToTsquery: query,
			Offset:    offset,
			Limit:     limit,
		})
		if err != nil && err != pgx.ErrNoRows {
			http.Error(w, "Search DB Error3", http.StatusInternalServerError)
			return
		}
	} else {
		pkgs, err = DBQueries.GetPackages(context.Background(), db.GetPackagesParams{
			Offset: offset,
			Limit:  limit,
		})
		if err != nil && err != pgx.ErrNoRows {
			http.Error(w, "DB Error4", http.StatusInternalServerError)
			return
		}
	}

	var cards []views.Card = make([]views.Card, len(pkgs))
	for i := range pkgs {
		cards[i].Name = pkgs[i].PkgName
		cards[i].URL = pkgs[i].Url
		cards[i].Description = pkgs[i].Description
	}
	q := r.URL.Query()
	q.Set("o", strconv.Itoa(int(offset-limit)))
	prevURL := q.Encode()
	q.Set("o", strconv.Itoa(int(offset+limit)))
	nextURL := q.Encode()
	w.WriteHeader(http.StatusOK)
	views.WriteIndexPage(w, cards, prevURL, nextURL)
}
