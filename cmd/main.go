package main

import (
	"fmt"
	"log"

	"github.com/jbnzi0/mixplorer/internal/music"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Mixplorer server")
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, token, err := music.GetToken()

	if err != nil {
		log.Fatalf("Couldn't get Spotify token: %v", err)
	}

	client := music.InitSpotifyClient(ctx, token)

	music.SearchNewReleases(client, ctx, "Travis Scott")
}
