package helpers

import "time"

func GetCurrentTimeVN() time.Time {
	location, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	currentTime := time.Now().In(location)

	return currentTime
}