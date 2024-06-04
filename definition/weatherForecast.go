package definition

type IWeatherForecast interface {
	GetWeatherInsight([]string, []string, []string) ([]*ConditionMetResult, int, error)
}
