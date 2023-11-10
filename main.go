package main

import (
	"context"
	"fmt"
	"intikom-interview/dal"
	"intikom-interview/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Initialize the database connection
	var err error
	db, err := gorm.Open(mysql.Open("root:tT£91F]y3[LV@tcp(127.0.0.1:3306)/intikom_test?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		panic("Failed to connect to the database")
	}

	dal.SetDefault(db)

	var clientID = "c93a80162b6637adb5ac"
	var clientSecret = "92878f248ba2e17dccfa21079d7b1a6845cc674e"
	// Set up the OAuth 2.0 config
	redirectURL := "http://localhost:8080/callback"

	oauthConfig := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
		//Scopes: []string{"scope1", "scope2"}, // Add required scopes
	}

	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		// Get the authorization URL
		authURL := oauthConfig.AuthCodeURL("state")
		fmt.Println("auth url = ", authURL)
		c.Redirect(http.StatusTemporaryRedirect, authURL)
	})

	// Route for handling the callback
	r.GET("/callback", func(c *gin.Context) {
		// Retrieve the authorization code from the callback URL query parameters
		authCode := c.Query("code")

		// Exchange the authorization code for an access token
		token, err := oauthConfig.Exchange(context.Background(), authCode)
		if err != nil {
			log.Fatal("Failed to exchange authorization code:", err)
		}

		// Send a response back to the browser
		c.JSON(http.StatusOK, gin.H{
			"access_token": token.AccessToken,
		})
	})

	routes.PublicRoutes(r)
	routes.PrivateRoutes(r, oauthConfig)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
