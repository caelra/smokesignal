package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/smoke_signal/internal/api"
	"github.com/smoke_signal/internal/github"
	"github.com/smoke_signal/internal/spotify"
	"github.com/smoke_signal/pkg/client"

	"github.com/caarlos0/env/v6"
	"github.com/spf13/cobra"
)

var (
	clientCfg client.Config

	rootCmd = &cobra.Command{
		Use:   "client",
		Short: "API Client",
		Long:  "API Client to run post deploy processing",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			log.Println("Starting Pre-run")
			if err := env.Parse(&clientCfg); err != nil {
				log.Fatalln("unable to load environment", err)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			githubClt := github.NewGithubClient(clientCfg.Github).GetGithubClient()
			spotifyClt := spotify.NewSpotifyClient(clientCfg.Spotify).GetToken().GetHttpClient().GetSpotifyClient()

			server := api.NewServer(&clientCfg, githubClt, spotifyClt)
			if err := server.Start(":3000"); err != nil {
				log.Println("server start failed: ", err)
			}
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {}
