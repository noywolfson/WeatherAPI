# WeatherAPI

Web service that will give the user a timeline of a weather insight - if a certain condition, based on the weather parameters, is met for a specific location.

To do so, I initialized a postgreSQL table with data from 3 given csv files. The files represent weather forecast data, each file represents a weather forecast of a different time (day/hour).  

The server will return either true or false for every timestamp, for two predefined conditions:
 
- veryHot - based on the condition temperature > 30
- rainyAndCold - based on the condition temperature < 10 AND precipitation > 0.5

## How to use this service

* Navigate to `https://weatherapi-jm68.onrender.com/weather/insight?condition={condition}&lat={lat}&lon={lon}`. replace the placeholders with the actual values.
  * for example: `https://weatherapi-jm68.onrender.com/weather/insight?condition=rainyAndCold&lat=24.5&lon=51.5`

- Alternatively, use curl from your terminal. 
    * for example: `curl -X GET "https://weatherapi-jm68.onrender.com/weather/insight?condition=rainyAndCold&lat=24.5&lon=51.5"`
