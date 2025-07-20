package http

import (
	"context"
	"log/slog"
	"net/http"
	"net/url"
)

type RequestURL struct {
	host string
	url  *url.URL
}

type requestURLType string

const requestURLKeyType = "httpRequestURL"

func RequestURLMiddleware(next http.Handler) http.Handler {
	slog.Info("Initializing Request middleware")
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(
				r.Context(),
				requestURLKeyType,
				RequestURL{
					r.Host,
					r.URL,
				},
			)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
