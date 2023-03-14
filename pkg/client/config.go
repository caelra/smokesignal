package client

import (
	"github.com/smoke_signal/internal/github"
	"github.com/smoke_signal/internal/spotify"
)

type Config struct {
    Env     string         `mapstructure:"env" env:"ENV"`
    Github  github.Config  `envPrefix:"GITHUB_"`
    Spotify spotify.Config `envPrefix:"SPOTIFY_"`
}

