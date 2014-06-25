package cc

import (
	"encoding/json"
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// AuthConfig holds basic OAuth information
type AuthConfig struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	VerCode        string `json:"verificationCode"`
}

// Cache writes an AuthConfig to disk
func (a *AuthConfig) Cache(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	b, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = f.Write(b)
	if err != nil {
		log.Println(err)
	}
	return f.Close()
}

// Load reads an AuthConfig from disk
func (a *AuthConfig) Load(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, a)

}

// Auth is a Requester that performes an OAuth request
type Auth struct {
	CacheFile string
	Config    *AuthConfig
}

// Get an http client which has been authorized by copy
func (a *Auth) Do() (*http.Client, error) {
	// If a cache file is given, read from the cache file
	if a.CacheFile != "" {
		tmpCache := &AuthConfig{}
		err := tmpCache.Load(a.CacheFile)
		if err != nil {
			log.Println("Could not load cache file with path: " + a.CacheFile)
			return nil, err
		}
		// replace empty entries with cached values
		if a.Config.ConsumerKey == "" {
			a.Config.ConsumerKey = tmpCache.ConsumerKey
		}
		if a.Config.ConsumerSecret == "" {
			a.Config.ConsumerSecret = tmpCache.ConsumerSecret
		}
		if a.Config.VerCode == "" {
			a.Config.VerCode = tmpCache.VerCode
		}
		// Cache the new settings
		err := a.Config.Cache(a.CacheFile)
		if err != nil {
			log.Println(err)
		}
	}
	c := oauth.NewConsumer(
		a.Config.ConsumerKey,
		a.Config.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.copy.com/oauth/request",
			AuthorizeTokenUrl: "https://www.copy.com/applications/authorize",
			AccessTokenUrl:    "https://api.copy.com/oauth/access",
		},
	)
	// If a verification code has not been supplied, request it from the browser and kill the app
	if a.Config.VerCode == "" {
		_, url, err := c.GetRequestTokenAndUrl("oob")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("(1) Go to: " + url)
		fmt.Println("(2) Grant access, you should get back a verification code.")
		fmt.Println("(3) Run the program again with command line argument -code $AUTHCODE")
		return nil, nil
	}
	// If a code is supplied, atempt to obtain an authorization
	accessToken, _, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	c.AuthorizeToken(accessToken, a.Config.VerCode)
	return nil, nil
}
