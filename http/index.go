package http

import (
	"net/http"
	"path/filepath"
	"selendis/g"
	"strings"

	"github.com/toolkits/file"
)


func configIndexRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			if !file.IsExist(filepath.Join(g.Root, "/static", r.URL.Path, "index.html")) {
				http.NotFound(w, r)
				return
			}
		}
		http.FileServer(http.Dir(filepath.Join(g.Root, "/static"))).ServeHTTP(w, r)
	})
}