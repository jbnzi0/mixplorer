package music

import (
	"context"
	"fmt"
	"log"
	"os"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func GetToken() (context.Context, *oauth2.Token, error) {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}

	token, err := config.Token(ctx)

	return ctx, token, err

}

func InitSpotifyClient(ctx context.Context, token *oauth2.Token) *spotify.Client {
	httpClient := spotifyauth.New().Client(ctx, token)
	return spotify.New(httpClient)
}

func SearchNewReleases(client *spotify.Client, ctx context.Context, artist string) {
	results, err := client.Search(ctx, artist, spotify.SearchTypeTrack|spotify.SearchTypeAlbum)
	if err != nil {
		log.Fatal(err)
	}

	if results.Albums != nil {
		fmt.Println("Albums:")
		for _, item := range results.Albums.Albums {
			fmt.Println("   ", item.Name)
		}
	}

	if results.Tracks != nil {
		fmt.Println("Tracks:")
		for _, item := range results.Tracks.Tracks {
			fmt.Println("   ", item.Name, item.Album.Name, item.Artists, item.Album.ReleaseDate)
		}
	}
}
