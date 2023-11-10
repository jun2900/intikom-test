package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func VerifyAccessToken(config *oauth2.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the access token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the token starts with "Bearer "
		parts := strings.Fields(authHeader)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Extract the access token
		accessToken := parts[1]

		// Create an OAuth2 token struct
		token := &oauth2.Token{
			AccessToken: accessToken,
			TokenType:   "Bearer",
		}
		// Create an HTTP client with the OAuth2 token
		httpClient := config.Client(context.Background(), token)

		// Use the httpClient to make authenticated requests
		resp, err := httpClient.Get("https://api.github.com/user")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to make request to github, details = %v", err)})
			c.Abort()
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Request failed with status: %s", resp.Status)})
			c.Abort()
			return
		}
	}
}
