package main

import (
	"fmt"
	"regexp"
)

var eventsRegexp = regexp.MustCompile(`\/events\/(\w+\-\w+\-\w+\-\w+\-\w+)`)

func main() {
	if eventsRegexp.MatchString("/events/a6e592bc-8627-4e13-b4a6-d7072864602a") {
		submatch := eventsRegexp.FindStringSubmatch("/events/a6e592bc-8627-4e13-b4a6-d7072864602a")
		uuid := submatch[1]
		fmt.Println(uuid)
	}
}
