package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return cors.AllowAll().Handler(next)
}
