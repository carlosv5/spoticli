package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/carlosv5/spoticli/pkg/credentials"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var userID = flag.String("user", "", "the Spotify user ID to look up")

func checkFlags() {
	if *userID == "" {
		flag.Usage()
		log.Fatalf("Error: missing user ID")
	}
}

func main() {
	flag.Parse()
	checkFlags()

	creds, err := credentials.Get()
	if err != nil {
		log.Fatalf("Wrong credentials")
	}

	config := &clientcredentials.Config{
		ClientID:     creds.Identifier,
		ClientSecret: creds.Secret,
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)
	user, err := client.GetUsersPublicProfile(spotify.ID(*userID))
	if err != nil {
		log.Fatalf("Error getting the information. Error: %v", err.Error())
	}

	fmt.Println("User ID:", user.ID)
	fmt.Println("Display name:", user.DisplayName)
	fmt.Println("Spotify URI:", string(user.URI))
	fmt.Println("Endpoint:", user.Endpoint)
	fmt.Println("Followers:", user.Followers.Count)
}
