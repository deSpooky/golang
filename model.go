package main

import "time"

type PostCard struct {
	ID       int           `json:"id"`
	Url      string        `json:"url"`
	Duration time.Duration `json:"duration"`
}