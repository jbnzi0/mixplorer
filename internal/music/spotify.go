package music

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type SearchResponse struct {
	Albums []spotify.SimpleAlbum
	Tracks []spotify.FullTrack
}

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

func SearchNewReleases(client *spotify.Client, ctx context.Context, artist string) SearchResponse {
	var response SearchResponse
	query := addFiltersToQuery(artist)
	results, err := client.Search(ctx, query, spotify.SearchTypeTrack|spotify.SearchTypeAlbum|spotify.SearchTypeArtist)

	if err != nil {
		log.Fatal(err)
	}

	if results.Albums != nil && len(results.Albums.Albums) > 0 {
		response.Albums = results.Albums.Albums
	}

	if results.Tracks != nil && len(results.Tracks.Tracks) > 0 {
		response.Tracks = results.Tracks.Tracks
	}

	return response
}

func addFiltersToQuery(query string) string {
	year, _, _ := time.Now().Date()
	return query + "%tag:new%" + "year:" + strconv.Itoa(year)
}
