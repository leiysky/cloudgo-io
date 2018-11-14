package handlers

import "net/http"

func StaticFileHandler(dir string) http.Handler {
	return http.FileServer(http.Dir(dir))
}
