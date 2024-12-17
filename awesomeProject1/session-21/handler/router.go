package handler

import "net/http"

func GetHandlers() *http.ServeMux {
	mux := http.NewServeMux()

	app := NewApplicationHandler()
	mux.HandleFunc("GET /api/v1/rates", app.GetRates)
	mux.HandleFunc("PUT /api/v1/rates", app.SetRate)
	return mux

}
