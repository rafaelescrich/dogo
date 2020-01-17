package main

import (
	"context"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
	"log"
	"os"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func main() {

	tkn := os.Getenv("DIGITAL_OCEAN_TOKEN")

	tokenSource := &TokenSource{
		AccessToken: tkn,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	log.Println(client)
}
