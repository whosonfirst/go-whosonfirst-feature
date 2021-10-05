package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_country = strings.NewReader(`{"properties":{"wof:country": "CA" }}`)

var missing_country = strings.NewReader(`{"properties":{ }}`)

func TestValidCountry(t *testing.T) {

	body, err := io.ReadAll(valid_country)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	country := Country(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if country != "CA" {
		t.Fatal("Invalid COUNTRY (valid)")
	}

}

func TestMissingCountry(t *testing.T) {

	body, err := io.ReadAll(missing_country)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	country := Country(body)

	if err != nil {
		t.Fatalf("Expect data (missing) failed")
	}

	if country != "XX" {
		t.Fatal("Invalid COUNTRY (missing)")
	}
}
