package controllers

import (
	"strconv"
	"strings"
	"github.com/akeeme/nfl-api-client/database"
	"github.com/akeeme/nfl-api-client/models"
	"github.com/gofiber/fiber/v2"
	_"fmt"
)

func GetAllStats(c *fiber.Ctx) error {
	var stats []models.Stat

	database.DB.Find(&stats)

	var statusDescription string

	// determine http status code
	if len(stats) == 0 {
		c.Status(404)
		statusDescription = "No stats found!"
	} else {
		c.Status(200)
		statusDescription = "Success!"
	}

	// send http status code with response
	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": statusDescription,
		"response": stats,
	})
}

func GetStatByPlayerId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		c.Next()
		return nil
	}

	stat := models.Stat {}

	database.DB.Find(&stat, id)

	var statusDescription string

	if stat.PLAYER_KEY == 0 {
		c.Status(404)
		statusDescription = "Player's Stat not found!"
		return c.JSON(fiber.Map{
			"status_code": c.Response().StatusCode(),
			"statusDescription": statusDescription,
			"response": nil,
		})
	} else {
		c.Status(200)
		statusDescription = "Success!"
	}

	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": statusDescription,
		"response": stat,
	})
}

func GetStatByPlayerName(c *fiber.Ctx) error {
	name := strings.Replace(c.Params("name"), "-", " ", -1)

	players := GetPlayersByName(name)

	var stat models.Stat

	if len(players) == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"status_code": c.Response().StatusCode(),
			"statusDescription": "Player not found!",
			"response": nil,
		})
	} else {
		for _, player := range players {
			database.DB.Find(&stat, player.PLAYER_KEY)
		}
		c.Status(200)
	}

	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": "Success!",
		"response": stat,
	})
	
}