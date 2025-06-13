package api

import (
	"encoding/json"
	"net/http"

	"github.com/manishkandari9/mcp-server/pkg/utils"
)

func FileTreeHandler(w http.ResponseWriter, r *http.Request) {
	files, err := utils.GetFileTree(".")
	if err != nil {
		http.Error(w, "Error while generating file tree", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // ✅ Good practice
	json.NewEncoder(w).Encode(files)
}
