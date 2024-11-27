package domain

import (
	"fmt"
	"projeto-docker/src/core"
	"projeto-docker/src/core/domain/errors"
	"strings"
	"time"
)

var logger = core.Logger()

type Model interface {
	IsValid() errors.Error
}

const UTCTimestampLayout = "2006-01-02 15:04:05 +0000 +0000"
const UTCDateLayout = "2006-01-02"

func BuildMapWithoutParentName(data map[string]interface{}, parentName string) map[string]interface{} {
	var newData = map[string]interface{}{}
	for k, v := range data {
		if strings.Contains(k, parentName) {
			newKey := strings.ReplaceAll(k, fmt.Sprintf("%s_", parentName), "")
			newData[newKey] = v
		}
	}
	return newData
}

func ParseUTCTimestampToDate(timestamp string) string {
	dt, err := time.Parse(UTCTimestampLayout, fmt.Sprint(timestamp))
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("Error when trying to parse timestamp to date: %s", err.Error()))
		return ""
	}
	return dt.Format(UTCDateLayout)
}

func ParseUTCTimestampToRFCNano(timestamp string) string {
	dt, err := time.Parse(UTCTimestampLayout, fmt.Sprint(timestamp))
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("Error when trying to parse UTC timestamp to RFC3339Nano: %s", err.Error()))
		return ""
	}
	return dt.Format(time.RFC3339Nano)
}

func ParseUTCTimestampToTime(timestamp string) *time.Time {
	dt, err := time.Parse(UTCTimestampLayout, fmt.Sprintf(timestamp))
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("Error when trying to parse UTC timestamp to time: %s", err.Error()))
		return nil
	}
	return &dt
}
