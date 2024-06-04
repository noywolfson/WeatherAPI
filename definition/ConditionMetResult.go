package definition

import "time"

type ConditionMetResult struct {
	ForecastTime time.Time
	ConditionMet bool
}
