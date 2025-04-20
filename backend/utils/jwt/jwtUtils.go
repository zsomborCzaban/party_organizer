package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zsomborCzaban/party_organizer/utils/env"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	GetIdFromJWTFunc = GetIdFromJWT // Default to real implementation
)

const ONE_DAY_IN_SECONDS = 86400
const ONE_HOUR_IN_SECONDS = 3600
const JWT_EXPIRATION_TIMEOUT_ENV_VAR_KEY = "JWT_EXPIRATION_TIMEOUT_KEY"
const JWT_SIGNING_KEY_ENV_VAR_KEY = "JWT_SIGNING_KEY"

const JWT_ISSUER_ENV_VAR_KEY = "JWT_ISSUER_KEY"
const JWT_ISSUER_DEFAULT_VALUE = "no_one"

func WithClaims(subject string, additionalClaims map[string]string, expirationTime int64) (*string, error) {
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

	secret := os.Getenv(JWT_SIGNING_KEY_ENV_VAR_KEY)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return nil, errors.New("token singing(lalala) failed: " + err.Error())
	}

	return &tokenString, err
}

// It's assumed that this is called after the jwt has been validated successfully
func GetIdFromJWT(bearer string) (uint, error) {
	tokenString := strings.Split(bearer, " ")

	token, err := jwt.Parse(tokenString[1], ParseToken)
	if err != nil || !token.Valid {
		return 0, errors.New("Error while parsing jwt")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Error while parsing the claims of jwt")
	}

	idUint64, err := strconv.ParseUint(claims["id"].(string), 10, 32)
	if err != nil {
		return 0, errors.New("Error while parsing the claims of the jwt")
	}

	return uint(idUint64), nil
}

func GetCanChangePasswordFromJWT(bearer string) (bool, error) {
	tokenString := strings.Split(bearer, " ")

	token, err := jwt.Parse(tokenString[1], ParseToken)
	if err != nil || !token.Valid {
		return false, errors.New("Error while parsing jwt")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, errors.New("Error while parsing the claims of jwt")
	}

	canChangeString, ok := claims["canChangePassword"].(string)
	if !ok {
		return false, nil
	}

	if canChangeString != "allowed" {
		return false, nil
	}

	return true, nil
}
