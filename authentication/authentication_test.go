package authentication_test

import (
	"savannahTest/authentication"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"savannahTest/config"
)

func TestGenerateSessionToken(t *testing.T) {

	email := "test@example.com"

	tokenString, err := authentication.GenerateSessionToken(email)

	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &authentication.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Configuration.GoogleClientSecret), nil
	})

	require.NoError(t, err)
	require.NotNil(t, token)
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	require.True(t, ok)
	assert.Equal(t, email, claims.Email)
	assert.WithinDuration(t, time.Now().Add(time.Hour*3), time.Unix(claims.ExpiresAt, 0), time.Minute)
}

func TestValidateSessionToken(t *testing.T) {

	email := "test@example.com"
	tokenString, err := authentication.GenerateSessionToken(email)
	require.NoError(t, err)

	valid, err := authentication.ValidateSessionToken(tokenString)
	require.NoError(t, err)
	assert.True(t, valid)

	invalidTokenString := "invalid.token.string"
	valid, err = authentication.ValidateSessionToken(invalidTokenString)
	assert.False(t, valid)
	assert.Error(t, err)

	expiredClaims := &authentication.JwtCustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-time.Hour).Unix(), // Token expired 1 hour ago
		},
	}
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
	expiredTokenString, err := expiredToken.SignedString([]byte(config.Configuration.GoogleClientSecret))
	require.NoError(t, err)

	valid, err = authentication.ValidateSessionToken(expiredTokenString)
	assert.False(t, valid)
	assert.Error(t, err)
}
