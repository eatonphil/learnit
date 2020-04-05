package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func getWords(inputFile string) []string {
	dict, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	words := []string{}
	for _, line := range strings.Split(string(dict), "\n") {
		if line == "" {
			continue
		}

		words = append(words, line)
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

	for i, word := range getWords(os.Args[1]) {
		// Set to 8am
		dateTime := time.Now().AddDate(0, 0, i).Truncate(time.Hour * 24).Add(time.Hour * 12)
		event := &calendar.Event{
			Summary: "[WOTD] " + word,
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
	}
}
