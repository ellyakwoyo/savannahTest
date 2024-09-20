package authentication

import (
	"github.com/golang-jwt/jwt"
	"savannahTest/config"
	"time"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

var secretKey = []byte(config.Configuration.GoogleClientSecret)

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateSessionToken(email string) (string, error) {
	claims := JwtCustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateSessionToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	_, ok := token.Claims.(*JwtCustomClaims)
	return ok && token.Valid, nil
}

func InitializeGoogleOAuth() {
	goth.UseProviders(
		google.New(config.Configuration.GoogleClientID, config.Configuration.GoogleClientSecret, config.Configuration.GoogleRedirectURI, "email", "profile"),
	)
}
