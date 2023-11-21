// SPDX-License-Identifier: Apache-2.0

// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

// The calendar-aware interval to use when bucketing. Possible values are: `minute`,
// `hour`, `day`, `week`, `month`, `quarter`, `year`.
type DateHistogramAggCalendarInterval uint8

const (
	DateHistogramAggCalendarIntervalMinute DateHistogramAggCalendarInterval = iota + 1
	DateHistogramAggCalendarIntervalHour
	DateHistogramAggCalendarIntervalDay
	DateHistogramAggCalendarIntervalWeek
	DateHistogramAggCalendarIntervalMonth
	DateHistogramAggCalendarIntervalQuarter
	DateHistogramAggCalendarIntervalYear
)

func (d DateHistogramAggCalendarInterval) String() string {
	switch d {
	default:
		return strconv.Itoa(int(d))
	case DateHistogramAggCalendarIntervalMinute:
		return "minute"
	case DateHistogramAggCalendarIntervalHour:
		return "hour"
	case DateHistogramAggCalendarIntervalDay:
		return "day"
	case DateHistogramAggCalendarIntervalWeek:
		return "week"
	case DateHistogramAggCalendarIntervalMonth:
		return "month"
	case DateHistogramAggCalendarIntervalQuarter:
		return "quarter"
	case DateHistogramAggCalendarIntervalYear:
		return "year"
	}
}

func (d DateHistogramAggCalendarInterval) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", d.String())), nil
}

func (d *DateHistogramAggCalendarInterval) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "minute":
		value := DateHistogramAggCalendarIntervalMinute
		*d = value
	case "hour":
		value := DateHistogramAggCalendarIntervalHour
		*d = value
	case "day":
		value := DateHistogramAggCalendarIntervalDay
		*d = value
	case "week":
		value := DateHistogramAggCalendarIntervalWeek
		*d = value
	case "month":
		value := DateHistogramAggCalendarIntervalMonth
		*d = value
	case "quarter":
		value := DateHistogramAggCalendarIntervalQuarter
		*d = value
	case "year":
		value := DateHistogramAggCalendarIntervalYear
		*d = value
	}
	return nil
}
