package utils

import "github.com/jpiontek/go-ip-api"

func GetGeolocation(ip string) *goip.Location {
	client := goip.NewClient()
	result, _ := client.GetLocationForIp(ip)
	return result
}
