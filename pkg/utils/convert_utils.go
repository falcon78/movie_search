package utils

import (
	"strconv"
	"strings"
	"time"
)

func Str2Int(value string) int64 {
	if integer, err := strconv.ParseInt(value, 10, 64); err != nil {
		return 0
	} else {
		return integer
	}
}

func Str2Float(value string) float64 {
	if float, err := strconv.ParseFloat(value, 64); err != nil {
		return 0
	} else {
		return float
	}
}

func Str2Bool(value string) bool {
	if strings.ToLower(value) == "true" {
		return true
	}
	return false
}

func Str2Date(value string) time.Time {
	if date, err := time.Parse("2006-01-02", value); err != nil {
		return time.Time{}
	} else {
		return date
	}
}
