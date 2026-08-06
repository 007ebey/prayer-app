package http

import nethttp "net/http"

func NewRouter() nethttp.Handler {
	mux := nethttp.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	return mux
}
