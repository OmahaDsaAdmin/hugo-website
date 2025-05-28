package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var ApiKey string
var ExportPath string

func main() {
	// Setup CLI args
	ApiKey = os.Args[1]
	ExportPath = os.Args[2]

	if ApiKey == "" || ExportPath == "" {
		fmt.Println("missing required args")
		return
	}

	// Setup http request
	url := "https://actionnetwork.org/api/v2/events/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("OSDI-API-Token", ApiKey)

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the response body into an event response
	var EventResponse EventResponse
	err = json.Unmarshal([]byte(body), &EventResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	// For each event, write the data to a markdown template
	events := EventResponse.Embedded.Events
	for i := 0; i < len(events); i++ {
		createAndFillTemplate(events[i])
	}

	fmt.Println("passed")
}

func createAndFillTemplate(e Event) {
	// Create template
	id := strings.Split(e.Identifiers[0], "_network:")[1]
	t, err := template.ParseFiles("template.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fix identifier, title, and description
	e.Identifiers = []string{id}
	e.Title = removeDoubleQuotes(e.Title)
	e.Description = stripHTMLTags(e.Description)

	// Check if a file already exists, if it does, skip it entirely so humans can control it
	newFilePath := ExportPath + "/" + id + ".md"
	fmt.Println(newFilePath)
	_, err = os.Stat(newFilePath)
	if err == nil {
		log.Println("skipped file: "+newFilePath, nil)
		return
	}

	// If no file exists
	f, err := os.Create(newFilePath)
	if err != nil {
		log.Println("create file: "+newFilePath, err)
		return
	}

	err = t.Execute(f, e)
	if err != nil {
		log.Print("execute: "+newFilePath, err)
		return
	}

	f.Close()
}

func stripHTMLTags(input string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(input, "")
}

func removeDoubleQuotes(input string) string {
	return strings.ReplaceAll(input, "\"", "")
}
