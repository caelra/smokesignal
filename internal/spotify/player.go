package spotify

import (
	"context"
	"log"

	"github.com/zmb3/spotify/v2"
)

func (c *Client) PlaySong(ctx context.Context, URIs []spotify.URI, deviceID spotify.ID) error {
	opts := &spotify.PlayOptions{URIs: URIs, DeviceID: &deviceID}

	if err := c.SpotifyClt.PlayOpt(ctx, opts); err != nil {
		log.Printf("error occurred trying to play a song: %s", err)
		return err
	}

	log.Printf("playing on device id: %s", deviceID)
	return nil
}

