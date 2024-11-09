package main

import (
	"github.com/rs/zerolog/log"
	"github.com/zsomborCzaban/party_organizer/configuration"
	"net/http"
)

func main() {
	log.Fatal().Err(http.ListenAndServe(":8080", configuration.DefaultApplicationSetup()))
}
