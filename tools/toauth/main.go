package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"time"
)

var LinuxDoEndpoint = oauth2.Endpoint{
	AuthURL:   "https://connect.linux.do/oauth2/authorize",
	TokenURL:  "https://connect.linux.do/oauth2/token",
	AuthStyle: oauth2.AuthStyleInHeader,
}

func main() {
	//ctx := context.Background()
	AuthCode()

	code := "9blgCB31PW8oMTLPNJP9NX9SKJ7qOjGT"
	tok, err := GetToken(code)
	fmt.Println(tok)
	fmt.Println(err)
}

func GetConf() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "hi3geJYfTotoiR5S62u3rh4W5tSeC5UG",
		ClientSecret: "VMPBVoAfOB5ojkGXRDEtzvDhRLENHpaN",
		//Scopes:       []string{"SCOPE1", "SCOPE2"},
		Endpoint: LinuxDoEndpoint,

		//RedirectURL: "http://127.0.0.1:8181/oauth2/redirectHere",
	}

	return conf
}

func AuthCode() {
	conf := GetConf()

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)
	fmt.Println()
}

func GetToken(code string) (*oauth2.Token, error) {
	conf := GetConf()
	//conf.RedirectURL = "http://127.0.0.1:8181/oauth2/redirectHere"
	ctx := context.Background()
	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 10 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	return tok, err
}
