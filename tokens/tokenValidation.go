package tokens

import (
	generalUtilities "SleekSpace/utilities/funcs/general"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func ValidateAccessToken(clientToken string) (claims *AccessTokenClaims, msg string) {
	token, err := jwt.ParseWithClaims(clientToken, &AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(generalUtilities.GetEnvVariables().TokenSecret), nil
	})
	if err != nil {
		msg = err.Error()
	}
	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok {
		msg = "token is invalid"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token has expired"
		return
	}
	return claims, msg
}
