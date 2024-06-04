# WeatherAPI

Web service that will give the user a timeline of a weather insight - if a certain condition, based on the weather parameters, is met for a specific location.

To do so, I initialized a postgreSQL table with data from 3 given csv files. The files represent weather forecast data, each file represents a weather forecast of a different time (day/hour).  

## How to use this service
### Requirements:
- Docker installed on your machine

### How to run:
* Clone the repository
* Navigate to the project directory in your terminal
* Run:

```bash
 docker compose up --build
```

* Navigate in your browser to `http://localhost:8080//weather/insight?condition={condition}&lat={lat}&lon={lon}`, and replace the placeholders with the actual values. 
  * for example: `http://localhost:8080/weather/insight?lon=51.5&lat=24.5&condition=veryHot`

  alternatively, use curl from your terminal. 
  * for example: `curl -X GET "http://localhost:8080/weather/insight?lon=51.5&lat=24.5&condition=veryHot"`
