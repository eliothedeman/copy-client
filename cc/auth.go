package cc

import (
	"fmt"
	"github.com/mrjones/oauth"
	"log"
	"net/http"
)

type AuthConfig struct {
	ConsumerKey, ConsumerSecret, VerCode string
}
type Auth struct {
	Config *AuthConfig
}

// Get an http client which has been authorized by copy
func (a *Auth) Do() (*http.Client, error) {
	c := oauth.NewConsumer(
		a.Config.ConsumerKey,
		a.Config.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.copy.com/oauth/request",
			AuthorizeTokenUrl: "https://www.copy.com/applications/authorize",
			AccessTokenUrl:    "https://api.copy.com/oauth/access",
		},
	)
	// If a verificatino code has not been supplied, request it from the browser and kill the app
	if a.Config.VerCode == "" {
		_, url, err := c.GetRequestTokenAndUrl("oob")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("(1) Go to: " + url)
		fmt.Println("(2) Grant access, you should get back a verification code.")
		fmt.Println("(3) Run the program again with command line argument -code $AUTHCODE")
		log.Fatal()
	}
	return nil, nil
}
