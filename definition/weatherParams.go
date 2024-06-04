package definition

import "time"

type WeatherParams struct {
	Longitude     float32
	Latitude      float32
	ForecastTime  time.Time
	Temperature   float32
	Precipitation float32
}
