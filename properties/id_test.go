package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_id = strings.NewReader(`{"properties":{"wof:id": 1234 }}`)

var invalid_id = strings.NewReader(`{"properties":{"wof:id": -5 }}`)

var missing_id = strings.NewReader(`{"properties":{ }}`)

func TestValidId(t *testing.T) {

	body, err := io.ReadAll(valid_id)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	id, err := Id(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if id != 1234 {
		t.Fatal("Invalid ID (valid)")
	}

}

func TestInValidId(t *testing.T) {

	body, err := io.ReadAll(invalid_id)

	if err != nil {
		t.Fatalf("Failed to read data (invalid), %v", err)
	}

	_, err = Id(body)

	if err == nil {
		t.Fatalf("Expect data (invalid) to fail")
	}
}

func TestMissingId(t *testing.T) {

	body, err := io.ReadAll(missing_id)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	_, err = Id(body)

	if err == nil {
		t.Fatalf("Expect data (missing) to fail")
	}
}
