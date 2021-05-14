package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var userID = flag.String("user", "", "the Spotify user ID to look up")

// Credentials of Spotify application
type Credentials struct {
	Identifier string `json:"SPOTIFY_ID"`
	Secret     string `json:"SPOTIFY_SECRET"`
}

func main() {
	flag.Parse()

	if *userID == "" {
		fmt.Fprintf(os.Stderr, "Error: missing user ID\n")
		flag.Usage()
		return
	}

	credentials := readCredentials()

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

func readCredentials() Credentials {
	jsonFile, err := os.Open("credentials/credentials.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened credentials.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var credentials Credentials
	json.Unmarshal(byteValue, &credentials)
	return credentials
}
