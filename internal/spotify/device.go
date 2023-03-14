package spotify

import (
	"context"
	"log"

	"github.com/zmb3/spotify/v2"
)

func (c *Client) AvaliableDevice(ctx context.Context) (spotify.ID, error) {
	devices, err := c.SpotifyClt.PlayerDevices(ctx)
	if err != nil {
		log.Printf("error trying to fetch avaliable player devices: %s", err)
		return "", err
	}

	var activeDeviceID spotify.ID = ""

	for _, device := range devices {
		if device.Active {
			activeDeviceID = device.ID
			break
		}
	}

	return activeDeviceID, nil
}

