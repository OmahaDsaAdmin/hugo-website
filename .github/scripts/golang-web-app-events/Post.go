package main

import (
	"time"
)

const TableNamePosts = "posts"

type EmbeddedPost struct {
	Status string `json:"Status"`
	Data   []Post `json:"Data"`
}

type Post struct {
	Id                int32     `json:"Id"`
	ActionNetworkGuid string    `json:"ActionNetworkGuid"`
	Title             string    `json:"Title"`
	Description       string    `json:"Description"`
	AuthorPublicName  string    `json:"AuthorPublicName"`
	EventStartDate    time.Time `json:"EventStartDate"`
	EventLink         string    `json:"EventLink"`
	EventImageLink    string    `json:"EventImageLink"`
	IsPublished       bool      `json:"IsPublished"`
	CreatedAt         time.Time `json:"CreatedAt"`
	UpdatedAt         time.Time `json:"UpdatedAt"`
	DeletedAt         time.Time `json:"DeletedAt"`
	CreatedBy         string    `json:"CreatedBy"`
	UpdatedBy         string    `json:"UpdatedBy"`
	DeletedBy         string    `json:"DeletedBy"`
	CreatedById       int32     `json:"CreatedById"`
	UpdatedById       int32     `json:"UpdatedById"`
	DeletedById       int32     `json:"DeletedById"`
	IsDeleted         int32     `json:"IsDeleted"`
}
