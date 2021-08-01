package middleware

import (
	"net/http"
)

func middleware(next http.Handler) http.Handler {
	return next
}
