package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/angusbean/enviro-check/app/models"
	"github.com/angusbean/enviro-check/platform/database"
	"github.com/gofiber/fiber/v2"
)

//RequestWeather
func RequestWeather(c *fiber.Ctx) error {

	coordinates := &models.Coord{}

	//Checking if data received from JSON body
	if err := c.BodyParser(coordinates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	closestCityID, err := db.LocateCity(coordinates.Lat, coordinates.Long)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	weatherModel, err := RetrieveWeather(closestCityID)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(weatherModel)
}

//RetrieveWeather returns the weather information based on the city ID from OpenWeather
func RetrieveWeather(closestCityID int) (models.WeatherUpdate, error) {
	var weatherModel models.WeatherUpdate

	API_KEY := os.Getenv("API_KEY")

	//Create API call
	APICall := "http://api.openweathermap.org/data/2.5/weather?id=" + strconv.Itoa(closestCityID) + "&appid=" + API_KEY

	//Create client & request
	client := &http.Client{}
	req, err := http.NewRequest("GET", APICall, nil)
	if err != nil {
		return weatherModel, err
	}

	//Add Request Headers and send
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return weatherModel, err
	}
	defer resp.Body.Close()

	//Read Response Body into Memory as bytes
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return weatherModel, err
	}

	//Unmarshal the bytes as json into the weather model
	json.Unmarshal(bodyBytes, &weatherModel)

	return weatherModel, nil
}
