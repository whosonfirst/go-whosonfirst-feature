package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_hierarchies = strings.NewReader(`{"properties":{"wof:hierarchy": [ {"foo_id": 123 }, { "bar_id": 456 } ]}}`)

var missing_hierarchies = strings.NewReader(`{"properties":{ }}`)

func TestValidHierarchies(t *testing.T) {

	body, err := io.ReadAll(valid_hierarchies)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	hierarchies := Hierarchies(body)

	if hierarchies == nil {
		t.Fatalf("Missing hierarchies")
	}

	if len(hierarchies) != 2 {
		t.Fatal("Invalid hierarchies, expected (2) results")
	}

	if hierarchies[0]["foo_id"] != 123 {
		t.Fatalf("Invalid hierarchies[0][\"foo_id\"] value")
	}

}

func TestMissingHierarchies(t *testing.T) {

	body, err := io.ReadAll(missing_hierarchies)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	hierarchies := Hierarchies(body)

	if hierarchies != nil {
		t.Fatalf("Hierarchies should be nil")
	}
}
