package license

import (
	"fmt"
	"testing"
)

func TestNewInfo(t *testing.T) {
	// Some values
	name, year := "Test", 1

	// Create an Info
	info := NewInfo(name, year)

	// Check the name
	if info.Name != name {
		t.Error("Expected name to be", name, "but was", info.Name)
	} //if

	// Check the year
	if info.Year != year {
		t.Error("Expected year to be", year, "but was", info.Year)
	} //if
} //TestNewInfo

func TestInfoReplace(t *testing.T) {
	// Some values
	name, year, lic := "Test", 1, "NAME-YEAR"

	// Create the expected
	expected := fmt.Sprintf("%s-%d", name, year)

	// Create the Info
	info := NewInfo(name, year)

	// Check replace
	if actual := info.replace(lic); actual != expected {
		t.Errorf("Expected the license to be \"%s\" but was \"%s\"", expected,
			actual)
	} //if
} //TestInfoReplace
