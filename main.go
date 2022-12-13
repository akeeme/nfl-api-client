package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/akeeme/nfl-api-client/database"
	"github.com/akeeme/nfl-api-client/routes"
	_"fmt"
)


func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)
	app.Listen(":3000")

	// fmt.Println(app)

	


	// // connect to db and get all players
	// database.Connect()

	// // get all players
	// var players []models.Player
	// database.DB.Find(&players)

	// for _, player := range players {
	// 	fmt.Println(player.Name)
	// }
}