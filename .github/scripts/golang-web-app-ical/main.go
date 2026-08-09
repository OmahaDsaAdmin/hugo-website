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

var ExportPath string
var ApiUrl string
var ApiKey string

func main() {
	ExportPath = os.Args[1]
	ApiUrl = os.Args[2]
	ApiKey = os.Args[3]

	// Setup http request
	url := ApiUrl
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Key", ApiKey)

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

	// Unmarshal the response body into an post response
	var ep EmbeddedPost
	err = json.Unmarshal([]byte(body), &ep)
	if err != nil {
		fmt.Println(err)
		return
	}

	posts := ep.Data
	if len(posts) == 0 {
		return
	}

	// For each event, write the data to a markdown template
	createAndFillTemplateWithList(posts)
}

func createAndFillTemplateWithList(posts []Post) {
	// Fix the data in the array
	newPosts := []Post{}

	for _, e := range posts {
		e.Title = removeDoubleQuotes(e.Title)
		e.Description = stripHTMLTags(e.Description)

		if e.ActionNetworkGuid != "" {
			newPosts = append(newPosts, e)
		}
	}

	// Create template
	t, err := template.ParseFiles("template.ics")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create the file, overwriting any old one
	newFilePath := ExportPath + "/" + "calendar_events.ical"
	fmt.Println(newFilePath)
	f, err := os.Create(newFilePath)
	if err != nil {
		log.Println("create file: "+newFilePath, err)
		return
	}

	err = t.Execute(f, newPosts)
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
