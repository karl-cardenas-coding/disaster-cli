package library

import (
	"testing"
)

func TestEventStrcuts(t *testing.T) {
	urlEvents := "https://eonet.sci.gsfc.nasa.gov/api/v3/events"
	var event EventResponse

	err := getJson(urlEvents, event)
	if err != nil {
		panic(err)
	}

	want := event.Events[0].Title

	for index, value := range event.Events {
		if index == 0 {
			got := value.Title
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		}

	}

}

func TestCategoriestrcuts(t *testing.T) {

	urlCategories := "https://eonet.sci.gsfc.nasa.gov/api/v3/categories/"
	CategoriesRes := new(CategoriesResponse)

	err := getJson(urlCategories, CategoriesRes)
	if err != nil {
		panic(err)
	}

	want := CategoriesRes.Categories[0].Title

	for index, value := range CategoriesRes.Categories {
		if index == 0 {
			got := value.Title
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		}

	}

}
