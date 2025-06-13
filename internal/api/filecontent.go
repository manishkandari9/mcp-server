package api

import (
	"io/ioutil"
	"net/http"
)

func FileContentHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("path")
	if file == "" {
		http.Error(w, "Missing 'path' query param", http.StatusBadRequest)
		return
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		http.Error(w, "File not found or can't be read", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
