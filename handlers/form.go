package handlers

import "net/http"

func FormHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	text := "<html><body><span>" + req.Form.Get("text1") + " " + req.Form.Get("text2") + "</span><br/><a href=\"/\">click to get back</a></body></html>"
	w.Write([]byte(text))
}
