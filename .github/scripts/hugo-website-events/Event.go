package main

import (
	"time"
)

type EventResponse struct {
	Embedded Embedded `json:"_embedded"`
}

type Embedded struct {
	Events []Event `json:"osdi:events"`
}

type Event struct {
	Identifiers  []string  `json:"identifiers"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Instructions string    `json:"instructions"`
	CreatedDate  time.Time `json:"created_date"`
	StartDate    time.Time `json:"start_date"`
	SponsorTitle string    `json:"action_network:sponsor.title"`
	Status       string    `json:"status"`
	Visibility   string    `json:"visibility"`
	EventLink    string    `json:"browser_url"`
	ImageLink    string    `json:"featured_image_url"`
}
