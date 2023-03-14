package api

import (
	"github.com/smoke_signal/internal/github"
	"github.com/smoke_signal/internal/spotify"
	"github.com/smoke_signal/pkg/client"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	router     *fiber.App
	githubClt  *github.Client
	spotifyClt *spotify.Client
	clientCfg  client.Config
}

func NewServer(clientCfg *client.Config, githubClt *github.Client, spotifyClt *spotify.Client) *Server {
	server := &Server{
		clientCfg:  *clientCfg,
		githubClt:  githubClt,
		spotifyClt: spotifyClt,
	}

	return server
}

func (server *Server) SetupRouter() {
	router := fiber.New()
	router.Use(logger.New())

	router.Post("/deploy", nil)

	server.router = router
}

func (server *Server) Router() *fiber.App {
	return server.router
}

func (server *Server) Start(port string) error {
	return server.router.Listen(port)
}
