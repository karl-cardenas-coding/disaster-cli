package library

import (
	"testing"
)

func TestEventStrcuts(t *testing.T) {
	urlEvents := "https://eonet.sci.gsfc.nasa.gov/api/v3/events"
	event := new(EventResponse)

	getJson(urlEvents, event)

	for _, value := range event.Events {
		got := value.Title
		want := "Wildfire - Vichuquen, Chile"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

}

func TestCategoriestrcuts(t *testing.T) {

	urlCategories := "https://eonet.sci.gsfc.nasa.gov/api/v3/categories/"
	CategoriesRes := new(CategoriesResponse)

	getJson(urlCategories, CategoriesRes)

	for _, value := range CategoriesRes.categories {
		got := value.Title
		want := "EONET Events: Drought"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

}
