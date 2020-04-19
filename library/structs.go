package library

import (
	"time"
)

type Response struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Events      []Events `json:"events"`
}
type Categories struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
type Sources struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
type Geometry struct {
	MagnitudeValue float64   `json:"magnitudeValue"`
	MagnitudeUnit  string    `json:"magnitudeUnit"`
	Date           time.Time `json:"date"`
	Type           string    `json:"type"`
	Coordinates    []float64 `json:"coordinates"`
}
type Events struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description interface{}   `json:"description"`
	Link        string        `json:"link"`
	Closed      interface{}   `json:"closed"`
	Categories  []Categories  `json:"categories"`
	Sources     []Sources     `json:"sources"`
	Geometry    []interface{} `json:"geometry"`
}
