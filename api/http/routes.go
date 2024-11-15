package http

import (
	"net/http"

	"github.com/jerhol/Hexel/api/http/handlers"
)

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)

	staticDir := http.Dir("web/static")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

	return mux
}
