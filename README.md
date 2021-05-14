# Spoticli
Spotify CLI for Go lang

# Installation and use
This is for now a first skeleton app in order to retrieve your user info.

To install:
1. To download vendors: `go mod vendor`
2. Fill your credentials in credentials/credentials.json
3. To build `go build cmd/spoticli/main.go`
4. To run with your userID as a flag `./main -user=YOUR_USER_ID`