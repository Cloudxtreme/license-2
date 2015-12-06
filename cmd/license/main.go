package main

import (
	"flag"
	"fmt"
	"os/user"
	"time"
)

var (
	defaultLicense = "MIT"
	defaultName    string
	defaultYear    int
)

func init() {
	// Determine the default year
	defaultYear = time.Now().Year()

	// Determine the default name
	if user, err := user.Current(); err != nil {
		// Panic
		panic(err)
	} else {
		// Set the name
		defaultName = user.Name
	} //if
} //init

func main() {
	// Create the flag for the name
	name := flag.String("name", defaultName, "The name of the copyright holder")

	// Create the flag for the year
	year := flag.Int("year", defaultYear, "The year of the copyright")

	// Create the flag for the license
	lic := flag.String("license", defaultLicense, "The license")

	switch *lic {
	case "MIT", "mit":
		fmt.Printf("Writing MIT license file for %s (c) %d\n", *name, *year)
	default:
		fmt.Printf("\"%s\" is not a valid license type.\n", *lic)
	} //switch
} //main
