package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zsomborCzaban/party_organizer/common/env"
	"time"
)

const ONE_DAY_IN_SECONDS = 86400
const JWT_EXPIRATION_TIMEOUT_ENV_VAR_KEY = "JWT_EXPIRATION_TIMEOUT_KEY"
const JWT_SIGNING_KEY_ENV_VAR_KEY = "JWT_SIGNING_KEY"

const JWT_ISSUER_ENV_VAR_KEY = "JWT_ISSUER_KEY"
const JWT_ISSUER_DEFAULT_VALUE = "ask peti about this"

func WithClaims(subject string, additionalClaims map[string]string) (*string, error) {
	expirationTimeout := env.GetEnvInt64(JWT_EXPIRATION_TIMEOUT_ENV_VAR_KEY, ONE_DAY_IN_SECONDS)

	//singingKeyString := env.GetEnvString(JWT_SIGNING_KEY_ENV_VAR_KEY, "")
	//fmt.Println(singingKeyString + "asd")
	//
	//if singingKeyString == "" {
	//	panic(fmt.Sprintf("%s environment variable not defined", JWT_SIGNING_KEY_ENV_VAR_KEY))
	//}
	//singingKey := []byte(singingKeyString)

	issuer := env.GetEnvString(JWT_ISSUER_ENV_VAR_KEY, JWT_ISSUER_DEFAULT_VALUE)

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, standardClaims)

	////for ES256 method
	//privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//fmt.Println(privateKey)
	//if err != nil {
	//	fmt.Println("Error generating key:", err)
	//	panic("lol")
	//}

	tokenString, err := token.SignedString([]byte("secret123"))

	if err != nil {
		return nil, errors.New("token singing(lalala) failed: " + err.Error())
	}

	return &tokenString, err
}
