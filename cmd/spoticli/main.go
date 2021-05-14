package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/carlosv5/spoticli/pkg/credentials"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var userID = flag.String("user", "", "the Spotify user ID to look up")

func main() {
	flag.Parse()

	if *userID == "" {
		fmt.Fprintf(os.Stderr, "Error: missing user ID\n")
		flag.Usage()
		return
	}

	credentials := credentials.Get()

	config := &clientcredentials.Config{
		ClientID:     credentials.Identifier,
		ClientSecret: credentials.Secret,
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)
	user, err := client.GetUsersPublicProfile(spotify.ID(*userID))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Println("User ID:", user.ID)
	fmt.Println("Display name:", user.DisplayName)
	fmt.Println("Spotify URI:", string(user.URI))
	fmt.Println("Endpoint:", user.Endpoint)
	fmt.Println("Followers:", user.Followers.Count)
}
