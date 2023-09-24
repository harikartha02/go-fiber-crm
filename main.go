package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/harikartha02/go-fiber-crm/database"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Post(NewLead)
	app.Get(GetLead)
	app.Delete(DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()

}
