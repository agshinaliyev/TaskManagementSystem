package server

import "net/http"

func Init() error {

	server := (&http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}).ListenAndServe()
	return server

}
