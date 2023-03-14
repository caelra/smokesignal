package entity

import (
	"github.com/zmb3/spotify/v2"
)

type DeployActonsRequest struct {
	URIs  []spotify.URI `json:"uris"`
	Repo  string        `json:"repo"`
	Owner string        `json:"owner"`
}
