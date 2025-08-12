package utils

import "time"

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func ValidateDateFormat(dateStr, format string) bool {
	_, err := time.Parse(format, dateStr)
	return err == nil
}
