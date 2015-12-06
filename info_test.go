package license

import "testing"

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
