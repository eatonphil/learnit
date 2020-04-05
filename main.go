package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func getWords() map[string]string {
	dict, err := ioutil.ReadFile("input.dict")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	words := map[string]string{}
	for _, line := range strings.Split(string(dict), "\n") {
		if line == "" {
			continue
		}

		sections := strings.SplitN(line, ":", 2)
		word := sections[0]
		trans := sections[1]
		words[word] = trans
	}

	return words
}

func main() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	i := 1
	for word, translation := range getWords() {
		// Set to 8am
		dateTime := time.Now().AddDate(0, 0, i).Truncate(time.Hour * 24).Add(time.Hour * 12)
		fmt.Println("Setting at", dateTime.Format(time.RFC3339), ", ", word, ":", translation)
		event := &calendar.Event{
			Summary: "[WOTD] " + word + ": " + translation,
			Start: &calendar.EventDateTime{
				DateTime: dateTime.Format(time.RFC3339),
				TimeZone: "America/New_York",
			},
			End: &calendar.EventDateTime{
				DateTime: dateTime.Add(time.Hour).Format(time.RFC3339),
				TimeZone: "America/New_York",
			},
		}

		calendarId := "primary"
		event, err = srv.Events.Insert(calendarId, event).Do()
		if err != nil {
			log.Fatalf("Unable to create event. %v\n", err)
		}
		fmt.Printf("Event created: %s\n", event.HtmlLink)
		i++
	}
}
