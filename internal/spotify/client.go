package spotify

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Client struct {
	SpotifyClt *spotify.Client
	Token      *oauth2.Token
	HttpClient *http.Client
	Config     Config
	Env        string
}

type ClientBuilder struct {
	Client *Client
}

var (
	authscopes = spotifyauth.WithScopes(
		spotifyauth.ScopeUserReadPrivate,
		spotifyauth.ScopeUserReadCurrentlyPlaying,
		spotifyauth.ScopeUserReadPlaybackState,
		spotifyauth.ScopeStreaming,
	)
)

func NewSpotifyClient(cfg Config) *ClientBuilder {
	return &ClientBuilder{Client: &Client{Config: cfg}}
}

func (cb *ClientBuilder) GetToken() *ClientBuilder {
	credsCfg := &clientcredentials.Config{
		ClientID:     cb.Client.Config.ID,
		ClientSecret: cb.Client.Config.Secret,
		TokenURL:     spotifyauth.TokenURL,
	}

	ctx := context.Background()
	token, err := credsCfg.Token(ctx)
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}

	cb.Client.Token = token

	return cb
}

func (cb *ClientBuilder) GetHttpClient() *ClientBuilder {
	ctx := context.Background()
	httpClient := spotifyauth.New(authscopes).Client(ctx, cb.Client.Token)

	cb.Client.HttpClient = httpClient

	return cb
}

func (cb *ClientBuilder) GetSpotifyClient() *Client {
	spotifyClt := spotify.New(cb.Client.HttpClient)
	cb.Client.SpotifyClt = spotifyClt

	return cb.Client
}
