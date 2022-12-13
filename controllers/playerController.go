package controllers

import (
	"strconv"
	"strings"
	"github.com/akeeme/nfl-api-client/database"
	"github.com/akeeme/nfl-api-client/models"
	"github.com/gofiber/fiber/v2"
	_"fmt"
)

func GetAllPlayers(c *fiber.Ctx) error {
	var players []models.Player
	database.DB.Find(&players)

	var statusDescription string

	// determine http status code
	if len(players) == 0 {
		c.Status(404)
		statusDescription = "No players found!"
	} else {
		c.Status(200)
		statusDescription = "Success!"
	}
	

	// send http status code with response
	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": statusDescription,
		"response": players,
	})
}

func GetPlayerById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		c.Next()
		return nil
	}

	player := models.Player {
		PLAYER_KEY: uint(id),
	}

	database.DB.Find(&player)

	var statusDescription string

	if player.Name == "" {
		c.Status(404)
		statusDescription = "Player not found!"
		return c.JSON(fiber.Map{
			"status_code": c.Response().StatusCode(),
			"statusDescription": statusDescription,
			"response": player,
		})
	}

	c.Status(200)
	statusDescription = "Success!"
	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": statusDescription,
		"response": player,
	})
}

func GetPlayerByName(c *fiber.Ctx) error {
	name := strings.Replace(c.Params("name"), "-", " ", -1)

	var players []models.Player
	database.DB.Where("name = ?", name).Find(&players)

	var statusDescription string

	// determine http status code
	if len(players) == 0 {
		c.Status(404)
		statusDescription = "No players found!"

		return c.JSON(fiber.Map{
			"status_code": c.Response().StatusCode(),
			"statusDescription": statusDescription,
			"response": nil,
		})
	} 
	
	c.Status(200)
	statusDescription = "Success!"


	// send http status code with response
	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": statusDescription,
		"response": players,
	})
}

func GetPlayersByName(name string) []models.Player {
	var players []models.Player
	database.DB.Where("name = ?", name).Find(&players)


	return players
}

func GetPlayersByTeamId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("teamId"))

	if err != nil {
		c.Next()
		return nil
	}

	var players []models.Player
	database.DB.Where("TEAM_KEY = ?", id).Find(&players)

	var statusDescription string

	// determine http status code
	if len(players) == 0 {
		c.Status(404)
		statusDescription = "No players found!"
	} else {
		c.Status(200)
		statusDescription = "Success!"
	}
	

	// send http status code with response
	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": statusDescription,
		"response": players,
	})
}

func GetPlayersByStatus(c *fiber.Ctx) error {
	status := c.Params("status")

	var players []models.Player
	database.DB.Where("status = ?", status).Find(&players)

	var statusDescription string

	// determine http status code
	if len(players) == 0 {
		c.Status(404)
		statusDescription = "No players found!"
	} else {
		c.Status(200)
		statusDescription = "Success!"
	}
	

	// send http status code with response
	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": statusDescription,
		"response": players,
	})
}

func GetPlayersByPosition(c *fiber.Ctx) error {
	position := c.Params("position")

	var players []models.Player
	database.DB.Where("position = ?", position).Find(&players)

	var statusDescription string

	// determine http status code
	if len(players) == 0 {
		c.Status(404)
		statusDescription = "No players found!"
	} else {
		c.Status(200)
		statusDescription = "Success!"
	}
	

	// send http status code with response
	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": statusDescription,
		"response": players,
	})
}


// NOT WORKING
func GetPlayersByTeamName(c *fiber.Ctx) error {
	name := strings.Replace(c.Params("teamName"), "-", " ", -1)

	teamKey := GetTeamKeyByName(name)

	var players []models.Player
	
	database.DB.Find(&players, "TEAM_KEY = ?", teamKey)

	var statusDescription string

	// determine http status code
	if len(players) == 0 {
		c.Status(404)
		statusDescription = "No players found!"
	} else {
		c.Status(200)
		statusDescription = "Success!"
	}

	// send http status code with response
	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": statusDescription,
		"response": players,
	})
}