package library

import (
	"encoding/json"
	"testing"
	// "fmt"
)

func TestEventStrcuts(t *testing.T) {

	inputJSON := "{\n\t\"id\": \"EONET_4637\",\n\t\"title\": \"Wildfire - Vichuquen, Chile\",\n\t\"description\": null,\n\t\"link\": \"https://eonet.sci.gsfc.nasa.gov/api/v3/events/EONET_4637\",\n\t\"closed\": null,\n\t\t\"categories\": [\n\t\t\t{\n\t\t\t\"id\": \"wildfires\",\n\t\t\t\"title\": \"Wildfires\"\n\t\t}\n\n\t\t\t\t\t\t],\n\t\t\"sources\": [\n\t\t\t{\n\t\t\t\"id\": \"PDC\",\n\t\t\t\"url\": \"http://emops.pdc.org/emops/?hazard_id=107743\"\n\t\t}\n\n\t\t\t\t\t\t\n\t],\n\t\t\"geometry\": [\n\t\t\t{\n\t\t\t\t\t\"magnitudeValue\": null,\n\t\t\t\"magnitudeUnit\": null,\n\t\t\t\"date\": \"2020-04-22T16:59:00Z\",\n\t\t\t\"type\": \"Point\", \n\t\t\t\t\t\"coordinates\": [ -72.02108, -34.84119 ]\n\t\t\t\t\t}\n\n\t\t\t\t]\n}"

	var response EventResponse

	err := json.Unmarshal([]byte(inputJSON), &response)
	if err != nil {
		t.Errorf("An error occured during unmarshal: %v\n Wanted: %v", response, "Pass")
	}

	for _, value := range response.Events {
		got := value.Title
		want := "Wildfire - Vichuquen, Chile"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

}

func TestCategoriestrcuts(t *testing.T) {

	inputJSON := "{\n\t\"title\": \"EONET Events: Drought\",\n\t\"description\": \"Long lasting absence of precipitation affecting agriculture and livestock, and the overall availability of food and water.\",\n\t\"link\": \"https://eonet.sci.gsfc.nasa.gov/api/v3/categories/drought\",\n\t\"events\": [\n\t]\n}"

	var response CategoriesResponse

	err := json.Unmarshal([]byte(inputJSON), &response)
	if err != nil {
		t.Errorf("An error occured during unmarshal: %v\n Wanted: %v", response, "Pass")
	}

	for _, value := range response.Categories {
		got := value.Title
		want := "EONET Events: Drought"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

}
