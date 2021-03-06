# enviro-check

To Do:
- Sort out who can create accounts
- Add Tests
- Add Docs and comments
- load Json file into db on request, should check for admin only
- Sort db migrations
- Use Swagger to generate docs
- Dockerise App


Enviro-Check is a webserver application that accepts JSON lat and long values to return live weather information from the closest weather station.

GO:
- built on Go (Golang) version 1.17

Uses:
- [OpenWeatherMap] for weather information (openweathermap.org)

Requires:
- Redis for as Key Value store
- Postgres for database non volatile memory storage 

Directions:
1) Create an acount and generate a valid OpenWeatherMap API token (https://home.openweathermap.org/users/sign_up)
2) Clone the repository
3) Update the values in .env-example and update name to .env
4) Ensure Redis is running with '$ redis-server'
5) Run with './run.sh'

Usage Example:

Send a 'Post' Request on 'localhost:5000/' as JSON, 
{
	"lat": -33.86,
	"long": 151.20
}

Application Returns:
{
    "name": "Ostrovnoy",
    "sys": {
        "country": "RU",
        "sunrise": 1629854503,
        "sunset": 1629912408
    },
    "coord": {
        "lat": 68.0531,
        "long": 0
    },
    "weather": [
        {
            "main": "Clouds",
            "description": "overcast clouds"
        }
    ],
    "main": {
        "temp": 282.16,
        "temp_min": 282.16,
        "temp_max": 282.16,
        "pressure": 1022,
        "humidity": 58
    },
    "visiblity": 0,
    "clouds": {
        "all": 87
    },
    "wind": {
        "speed": 4.79,
        "deg": 75,
        "gust": 3.79
    },
    "cod": 200
}

Performance:
- Calculates and returns weather information as JSON on average 1.2 seconds with 50up/20down connection speeds

Roadmap:
- Add authentication via JWT
- Use db for authentication 
- Improve JSON error responses
- Run as Docker for cloud deploy
- Add stream via gRPC

Minor ToDo's:
- Create Redis client as application variable (remove duplication of code)

Other Ideas:
- add density indicators (traffic? 4G coverage? housing data?)
- ArcGIS urban density calculators
- add low altitude airspace information (flight paths? ADS-B information)