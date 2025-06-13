package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rs/cors"

    "github.com/manishkandari9/mcp-server/internal/api"
    // "github.com/manishkandari9/mcp-server/internal/middleware"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to MCP Server 🚀"))
    })

    router.HandleFunc("/filetree", api.FileTreeHandler).Methods("GET")

    handler := cors.Default().Handler(router)

    log.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}

