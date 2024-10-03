package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zsomborCzaban/party_organizer/common/utils"
	"time"
)

const ONE_DAY_IN_SECONDS = 86400
const JWT_EXPIRATION_TIMEOUT_ENV_VAR_KEY = "JWT_EXPIRATION_TIMEOUT_KEY"
const JWT_SINGING_KEY_ENV_VAR_KEY = "JWT_SINGING_KEY"

const JWT_ISSUER_ENV_VAR_KEY = "JWT_ISSUER_KEY"
const JWT_ISSUER_DEFAULT_VALUE = "ask peti about this"

func WithClaims(subject string, additionalClaims map[string]string) (*string, error) {
	expirationTimeout := utils.GetEnvInt64(JWT_EXPIRATION_TIMEOUT_ENV_VAR_KEY, ONE_DAY_IN_SECONDS)

	singingKeyString := utils.GetEnvString(JWT_SINGING_KEY_ENV_VAR_KEY, "")

	if singingKeyString == "" {
		panic(fmt.Sprintf("%s environment variable not defined", JWT_SINGING_KEY_ENV_VAR_KEY))
	}
	singingKey := []byte(singingKeyString)

	issuer := utils.GetEnvString(JWT_ISSUER_ENV_VAR_KEY, JWT_ISSUER_DEFAULT_VALUE)

	now := time.Now()

	standardClaims := jwt.MapClaims{
		"iss": issuer,
		"sub": subject,
		"exp": now.Unix() + expirationTimeout,
		"nbf": now.Unix(),
		"iat": now.Unix(),
	}

	for k, v := range additionalClaims {
		standardClaims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, standardClaims)

	tokenString, err := token.SignedString(singingKey)

	if err != nil {
		return nil, errors.New("token singing(lalala) failed: " + err.Error())
	}

	return &tokenString, err
}
