package configuration

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	for _, envKey := range checkEnvKeyOnSetup {
		_, exists := os.LookupEnv(envKey)
		if !exists {
			log.Fatal().Msg("environment key missing: " + envKey)
		}
	}
}
