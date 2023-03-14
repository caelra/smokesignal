package api

import (
	"github.com/smoke_signal/internal/api/entity"

	"github.com/gofiber/fiber/v2"
)

func (server *Server) DeployActions(fiberctx *fiber.Ctx) {
	ctx := fiberctx.Context()

	var req *entity.DeployActonsRequest
	if err := fiberctx.BodyParser(&req); err != nil {
		fiberctx.Status(400).JSON(fiber.Map{"error": err})
		return
	}

	deviceID, err := server.spotifyClt.AvaliableDevice(ctx)
	if err != nil {
		fiberctx.Status(401).JSON(fiber.Map{"error": err})
		return
	}

	if err := server.spotifyClt.PlaySong(ctx, req.URIs, deviceID); err != nil {
		fiberctx.Status(401).JSON(fiber.Map{"error": err})
		return
	}

	fiberctx.Status(200).JSON(fiber.Map{"device_id": deviceID})
}
