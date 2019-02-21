package FKGoKylin

import "time"

type TimePair struct{
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (tp TimePair) IsZero() bool {
	return tp.StartTime.IsZero() && tp.EndTime.IsZero()
}
