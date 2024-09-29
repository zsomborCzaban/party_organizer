package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	log.Fatal().Err(http.ListenAndServe(":8080", router))
}
