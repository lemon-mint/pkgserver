package main

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/lemon-mint/pkgserver/db"
	"github.com/lemon-mint/pkgserver/views"
)

//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate go run github.com/valyala/quicktemplate/qtc -dir=views
//go:generate go mod tidy

type PackagesHandler struct{}

func (*PackagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pkg, err := DBQueries.GetPackageWithName(context.Background(), r.Host+r.URL.Path)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}
		http.Error(w, "DB Error0", http.StatusInternalServerError)
		return
	}

	switch pkg.PkgType {
	case db.PkgtypeGo:
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		views.WriteGoPackage(w, pkg.PkgName, string(pkg.Vcs), pkg.Url)
	default:
		http.Error(w, "Package Error1", http.StatusServiceUnavailable)
		return
	}
}
