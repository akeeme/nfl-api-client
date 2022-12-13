package routes

import (
	_"strconv"

	"github.com/akeeme/nfl-api-client/controllers"
	"github.com/gofiber/fiber/v2"
)


func Setup(app *fiber.App) {

	

	// teams
	teams := app.Group("/api/teams")
	teams.Get("/", controllers.GetAllTeams)
	teams.Get("/:id", controllers.GetTeamById)
	teams.Get("/:name", controllers.GetTeamByName)
	teams.Get("/:division", controllers.GetTeamsByDivision)
	teams.Put("/:id", controllers.PutTeam)
	teams.Delete("/:id", controllers.DeleteTeamById)
	teams.Post("/", controllers.PostTeam)

	// players
	players := app.Group("/api/players")
	players.Get("/", controllers.GetAllPlayers)
	players.Get("/:id", controllers.GetPlayerById)
	players.Get("/:name", controllers.GetPlayerByName)
	players.Get("/teamid/:teamId", controllers.GetPlayersByTeamId)
	players.Get("/status/:status", controllers.GetPlayersByStatus)
	players.Get("/position/:position", controllers.GetPlayersByPosition)
	players.Get("/team/:teamName", controllers.GetPlayersByTeamName)


	stats := app.Group("/api/stats")
	stats.Get("/", controllers.GetAllStats)
	stats.Get("/:id", controllers.GetStatByPlayerId)
	stats.Get("/:name", controllers.GetStatByPlayerName)


	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"status_code": c.Response().StatusCode(),
			"statusDescription": "Route not found!",
			"response": nil,
		})
	})

	
}