package handlers

import "net/http"

func UnknownHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`
	500 INTERNAL ERROR
	`))
}
