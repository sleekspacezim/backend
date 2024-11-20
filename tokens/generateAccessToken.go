package tokens

import (
	generalUtilities "SleekSpace/utilities/funcs/general"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AccessTokenClaims struct {
	GivenName string
	Email     string
	Id        string
	jwt.StandardClaims
}

func GenerateAccessToken(givenName string, email string, id int) string {
	claims := &AccessTokenClaims{
		GivenName: givenName,
		Email:     email,
		Id:        string(rune(id)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().AddDate(10, 0, 0).Unix(),
			Id:        string(rune(id)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(generalUtilities.GetEnvVariables().TokenSecret))
	if err != nil {
		println("error", err.Error())
		return "failed"
	} else {
		return tokenString
	}

}
