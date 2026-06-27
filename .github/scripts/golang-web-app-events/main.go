package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode"
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

	// For each post, write the data to a markdown template
	for i := 0; i < len(posts); i++ {
		createAndFillTemplate(posts[i], ExportPath, ApiKey)
	}
}

func createAndFillTemplate(p Post, ep string, k string) {
	// Get the id
	id := p.ActionNetworkGuid
	if p.ActionNetworkGuid == "" {
		id = strconv.Itoa(int(p.Id)) + "-" + toKebabCase(p.Title)
	}

	// Create template
	t, err := template.ParseFiles("template.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fix identifier, title, and description
	p.Title = removeDoubleQuotes(p.Title)
	p.Description = stripHTMLTags(p.Description)

	postPath := "posts"
	if p.ActionNetworkGuid != "" {
		postPath = "events"
	}

	// Check if a file already exists, if it does, skip it entirely so humans can control it
	newFilePath := ep + "/" + postPath + "/" + id + ".md"

	fmt.Println(newFilePath)
	_, err = os.Stat(newFilePath)
	if err == nil {
		os.Remove(newFilePath)
		log.Println("removing old file: "+newFilePath, nil)
	}

	// If no file exists
	f, err := os.Create(newFilePath)
	if err != nil {
		log.Println("create file: "+newFilePath, err)
		return
	}

	err = t.Execute(f, p)
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

func toKebabCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}

	var b strings.Builder
	b.Grow(len(s))

	prevHyphen := false

	for _, r := range strings.ToLower(s) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			prevHyphen = false
		default:
			// treat any non-alnum as a separator
			if !prevHyphen {
				b.WriteByte('-')
				prevHyphen = true
			}
		}
	}

	out := b.String()
	return strings.Trim(out, "-")
}
