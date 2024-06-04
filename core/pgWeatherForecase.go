package core

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tomorrow-code-challenge/backend-wdata-code-challenge.git/config"
	"github.com/tomorrow-code-challenge/backend-wdata-code-challenge.git/definition"
	"net/http"
	"strconv"
)

const (
	veryHot      string = "veryHot"
	rainyAndCold string = "rainyAndCold"
)

type PgWeatherForecast struct {
	pgDB *sql.DB
}

func NewPgWeatherForecast(pgDB *sql.DB) *PgWeatherForecast {
	return &PgWeatherForecast{pgDB: pgDB}
}

func (pg *PgWeatherForecast) GetWeatherInsight(conditionParam []string, latParam []string, lonParam []string) ([]*definition.ConditionMetResult, int, error) {
	floatLon, floatLat, condition, err := convertParamsAndValidate(conditionParam, lonParam, latParam)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE longitude = $1 AND latitude = $2", config.Static.DbName)
	rows, err := pg.pgDB.Query(query, floatLon, floatLat)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var result []*definition.ConditionMetResult
	for rows.Next() {
		row := &definition.WeatherParams{}
		err := rows.Scan(&row.Longitude, &row.Latitude, &row.ForecastTime, &row.Temperature, &row.Precipitation)
		if err != nil {
			logrus.WithError(err).Error()
			continue
		}

		conditionMetRes := checkConditionMet(row, condition)
		result = append(result, conditionMetRes)
	}
	return result, http.StatusOK, nil
}

func checkConditionMet(row *definition.WeatherParams, condition string) *definition.ConditionMetResult {
	res := &definition.ConditionMetResult{
		ForecastTime: row.ForecastTime,
		ConditionMet: false,
	}
	switch condition {
	case veryHot:
		if row.Temperature > 40 {
			res.ConditionMet = true
		}
	case rainyAndCold:
		if row.Temperature < 10 && row.Precipitation > 0.5 {
			res.ConditionMet = true
		}
	default:
		return nil
	}
	return res
}

func convertParamsAndValidate(conditionParam []string, lonParam []string, latParam []string) (float64, float64, string, error) {
	if len(conditionParam) == 0 || len(latParam) == 0 || len(lonParam) == 0 {
		return 0, 0, "", errors.New("condition, lon and lat are required")
	}

	condition := conditionParam[0]
	lon := lonParam[0]
	lat := latParam[0]

	var floatLat, floatLon float64
	var err error
	if floatLon, err = strconv.ParseFloat(lon, 64); err != nil {
		return 0, 0, "", errors.New("lon input is not a valid float")
	}
	if floatLat, err = strconv.ParseFloat(lat, 64); err != nil {
		return 0, 0, "", errors.New("lat input is not a valid float")
	}
	if condition == "" {
		return 0, 0, "", errors.New("condition can't be empty")
	}
	if condition != veryHot && condition != rainyAndCold {
		return 0, 0, "", errors.New("condition must be one of veryHot or rainyAndCold")
	}
	return floatLon, floatLat, condition, nil
}
