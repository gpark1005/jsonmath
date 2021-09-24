package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewServer() *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/data", PostNewData).Methods("POST")
	r.HandleFunc("/data", UpdateData).Methods("PUT")
	svr := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}

	return svr
}
