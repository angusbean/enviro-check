package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/angusbean/enviro-check/app/models"
	"github.com/angusbean/enviro-check/platform/database"
	"github.com/gofiber/fiber/v2"
)

func UpdateCities(c *fiber.Ctx) error {

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Delete all rows in cities table
	db.ClearCitiesTable()

	// Read JSON file into struct
	jsonFile, err := os.Open("platform/openweather/city.list.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	//Read opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//Initialise City array
	var cityList models.Cities

	//Unmarshal byteArray into cities struct
	json.Unmarshal(byteValue, &cityList)

	// Insert struct into db to populate db with list of cities
	db.InsertCityInfo(cityList)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Successfully updated cities table in database",
	})

}
