package api

import (
	"io"
	"net/http"
)

func GitHubHandler(w http.ResponseWriter, r *http.Request) {
	repo := r.URL.Query().Get("repo")
	if repo == "" {
		http.Error(w, "Missing 'repo' query param", http.StatusBadRequest)
		return
	}

	resp, err := http.Get("https://api.github.com/repos/" + repo)
	if err != nil {
		http.Error(w, "GitHub API call failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}
