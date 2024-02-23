package server

import (
	"log/slog"
	"net/http"

	"backend/app/middleware/loggermw"
)

func New(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	addRoutes(mux, logger)

	var handler http.Handler = mux
	handler = loggermw.New(logger, handler)

	return handler
}
