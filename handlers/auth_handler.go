package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
	"net/http"
	"savannahTest/authentication"
)

// Home godoc
// @Summary Show Home page with login option
// @Description Displays a simple HTML page with a link to log in with Google.
// @Tags Authentication
// @Produce html
// @Success 200 {string} string "HTML content"
// @Router / [get]
func Home(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/app/v1/login")
}

// Login godoc
// @Summary Start Google OAuth2 login
// @Description Redirects the user to Google's OAuth2 login page to begin authentication.
// @Tags Authentication
// @Produce json
// @Param provider query string true "OAuth provider" default(google)
// @Success 302 "Redirects to Google login page"
// @Router /login [get]
func Login(c echo.Context) error {

	q := c.Request().URL.Query()
	q.Add("provider", "google")
	c.Request().URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(c.Response(), c.Request())
	return nil
}

// HandleGoogleCallback godoc
// @Summary Google OAuth2 callback handler
// @Description Handles the callback from Google after authentication. Creates a session token and sets it in a cookie.
// @Tags Authentication
// @Produce json
// @Success 200 {object} map[string]interface{} "Authentication success with user and session token"
// @Failure 400 {string} string "Authentication failed"
// @Failure 500 {string} string "Session generation failed"
// @Router /auth/callback [get]
func HandleGoogleCallback(c echo.Context) error {

	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.String(http.StatusBadRequest, "Authentication failed: "+err.Error())
	}

	sessionToken, err := authentication.GenerateSessionToken(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate session"})
	}

	c.SetCookie(&http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":          user,
		"session_token": sessionToken,
	})
}
