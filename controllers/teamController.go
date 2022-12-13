package controllers

import (
	"strconv"
	"strings"
	"github.com/akeeme/nfl-api-client/database"
	"github.com/akeeme/nfl-api-client/models"
	"github.com/gofiber/fiber/v2"
	_"fmt"
)

func GetAllTeams(c *fiber.Ctx) error {
	var teams []models.Team
	database.DB.Find(&teams)

	var statusDescription string

	// determine http status code
	if len(teams) == 0 {
		c.Status(404)
		statusDescription = "No teams found!"
	} else {
		c.Status(200)
		statusDescription = "Success!"
	}
	

	// send http status code with response
	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": statusDescription,
		"response": teams,
	})
	
}


func GetTeamById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		c.Next()
		return nil
	}


	team := models.Team {
		TEAM_KEY: uint(id),
	}


	database.DB.Find(&team)

	var statusDescription string

	if team.Established == 0 {
		c.Status(404)
		statusDescription = "Team not found!"
		return c.JSON(fiber.Map{
			"status_code": c.Response().StatusCode(),
			"statusDescription": statusDescription,
			"response": nil,
		})
	} 

	c.Status(200)
	statusDescription = "Success!"
	
	return c.JSON(fiber.Map{
	"status_code": c.Response().StatusCode(),
	"statusDescription": statusDescription,
	"response": team,
	})
	
}

func GetTeamByName(c *fiber.Ctx) error {	
	teamName := strings.Replace(c.Params("name"), "-", " ", -1)

	team := models.Team {
		Name: teamName,
	}

	database.DB.Where("name = ?", teamName).Find(&team)

	var statusDescription string

	if team.Established == 0 {
		c.Next()
		return nil
	} 

	c.Status(200)
	statusDescription = "Success!"
	
	return c.JSON(fiber.Map{
	"status_code": c.Response().StatusCode(),
	"statusDescription": statusDescription,
	"response": team,
	"teamName": teamName,
	})
}

func GetTeamKeyByName(name string) uint {
	team := models.Team {
		Name: name,
	}

	database.DB.Where("name = ?", name).Find(&team)

	return team.TEAM_KEY
}

func GetTeamsByDivision(c *fiber.Ctx) error {
	division := strings.Replace(c.Params("division"), "-", " ", -1)

	var teams []models.Team

	database.DB.Where("division = ?", division).Find(&teams)

	var statusDescription string

	// determine http status code
	if len(teams) == 0 {
		c.Status(404)
		statusDescription = "No teams found!"
	} else {
		c.Status(200)
		statusDescription = "Success!"
	}
	

	// send http status code with response
	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": statusDescription,
		"response": teams,
	})
}

func PostTeam(c *fiber.Ctx) error {
	var team models.Team

	if err := c.BodyParser(&team); err != nil {
		return err
	}

	database.DB.Create(&team)

	if database.DB.Find(&team).RowsAffected == 0 {
		c.Status(500)
		return c.JSON(fiber.Map{
			"status_code": c.Response().StatusCode(),
			"statusDescription": "Error creating team!",
			"response": nil,
		})
	} else {
		c.Status(201)
		return c.JSON(fiber.Map{
			"status_code": c.Response().StatusCode(),
			"statusDescription": "Success!",
			"response": team,
		})
	}

}

func PutTeam(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	team := models.Team {
		TEAM_KEY: uint(id),
	}

	if err := c.BodyParser(&team); err != nil {
		return err
	}

	database.DB.Model(&team).Updates(team)

	database.DB.Find(&team)

	return c.JSON(fiber.Map{
		"status_code": c.Response().StatusCode(),
		"statusDescription": "Success!",
		"response": team,
	})
	// var team models.Team

	// if err := c.BodyParser(&team); err != nil {
	// 	return err
	// }

	// database.DB.Model(&team).Updates(team)

	// return c.JSON(team)
}

func DeleteTeamById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return nil
	}

	team := models.Team {
		TEAM_KEY: uint(id),
	}

	

	// if team exists in db then delete
	if database.DB.Find(&team).RowsAffected != 0 {
		database.DB.Delete(&team)
		c.Status(200)
		return c.JSON(fiber.Map{
			"status_code": c.Response().StatusCode(),
			"statusDescription": "Successfully Deleted!",
			"response": team,
		})
	} else {
		c.Status(404)
		return c.JSON(fiber.Map{
			"status_code": c.Response().StatusCode(),
			"statusDescription": "Team with id " + c.Params("id") + " not found!",
			"response": nil,
		})
	}
}
