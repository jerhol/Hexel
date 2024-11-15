package handlers

import (
	"net/http"
	"path/filepath"
)

// HomeHandler serves the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("web", "static", "index.html"))
}
