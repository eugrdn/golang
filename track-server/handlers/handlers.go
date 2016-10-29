package handlers

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

//Handler handle base requests
func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if len(path) == 0 {
		writePath(w, path)
		return
	}
	url, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		writePath(w, path)
		return
	}
	http.Redirect(w, r, "http://"+string(url), http.StatusFound)

}

//HandlerIcon handle favicon requests
func HandlerIcon(w http.ResponseWriter, r *http.Request) {
	// nothing doing
}

func writePath(w http.ResponseWriter, path string) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Hi", "Your path /"+path)
}
