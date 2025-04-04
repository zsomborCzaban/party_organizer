package configuration

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/zsomborCzaban/party_organizer/common"
	"os"
)

var checkEnvKeyOnSetup = []string{common.JWT_SINGING_KEY_ENV_KEY, common.AWS_ACCESS_KEY_ID_ENV_KEY, common.AWS_SECRET_ACCESS_KEY_ENV_KEY, common.AWS_REGION_ENV_KEY, common.AWS_BUCKET_NAME_ENV_KEY}

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
