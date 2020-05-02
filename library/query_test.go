package library

import (
	"testing"
)

func TestQueryEventsAPI(t *testing.T) {
	want := "Greater than 0"
	got := QueryEventAPI("")

	if len(got.Events) <= 0 {
		t.Errorf("Query events failed: %v \nwant: %v", got, want)
	}
}

func TestQueryCategoriesAPI(t *testing.T) {

	want := "Greater than 0"
	got := QueryCategoriesAPI("", "")

	if len(got.Categories) <= 0 {
		t.Errorf("Query Categories failed: %v \nwant: %v", got, want)
	}
}
