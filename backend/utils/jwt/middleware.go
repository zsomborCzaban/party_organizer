package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zsomborCzaban/party_organizer/utils/api"
	"net/http"
	"os"
	"strings"
)

func ParseToken(token *jwt.Token) (interface{}, error) {
	secret := os.Getenv(JWT_SIGNING_KEY_ENV_VAR_KEY)

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("jwt was signed with invalid method")
	}
	return []byte(secret), nil
}

func ValidateJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bearer := r.Header.Get("Authorization")

		tokenString := strings.Split(bearer, " ")
		if len(tokenString) < 2 {
			resp := api.ErrorBadRequest("Authorization missing from header")
			resp.Send(w)
			return
		}

		token, err := jwt.Parse(tokenString[1], ParseToken)
		if err != nil || !token.Valid {
			resp := api.ErrorUnauthorized("invalid jwt")
			resp.Send(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
