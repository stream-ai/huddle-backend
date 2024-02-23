package server

import (
	"io"
	"log"
	"net/http"
)

func handleHealthZ() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("health check\n")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "OK")
	}
}
