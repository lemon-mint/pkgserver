package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v4"
)

//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate go run github.com/valyala/quicktemplate/qtc -dir=views
//go:generate go mod tidy

type PackagesHandler struct{}

func (*PackagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pkg_name := strings.TrimPrefix(r.URL.Path, "/")
	pkg, err := DBQueries.GetPackageWithName(context.Background(), pkg_name)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}
		http.Error(w, "DB Error0", http.StatusInternalServerError)
		return
	}

	switch pkg.PkgType {
	default:
		http.Error(w, "DB Error1", http.StatusInternalServerError)
		return
	}
}
