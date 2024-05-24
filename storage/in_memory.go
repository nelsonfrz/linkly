package storage

import "fmt"

var urls map[string]string = make(map[string]string)
var counter int = 0

func SaveUrl(url string) string {
	counter++
	id := fmt.Sprintf("%d", counter)
	urls[id] = url
	return id
}

func GetUrl(id string) string {
	url := urls[id]
	return url
}
