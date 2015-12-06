package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/antizealot1337/license"
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

	// Parse the flags
	flag.Parse()

	var licType license.Type

	switch strings.ToLower(*lic) {
	case "mit":
		// Set the license type
		licType = license.MIT
	case "bsd":
		// Set the license type
		licType = license.BSD
	default:
		fmt.Printf("\"%s\" is not a valid license type.\n", *lic)
		os.Exit(-1)
	} //switch

	// Create the info
	info := license.NewInfo(*name, *year)

	// Create the file
	file, err := os.Create("LICENSE")

	// Check if there was an error
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	} //if

	// Make sure the file is closed
	defer file.Close()

	// Write the license
	license.Write(licType, info, file)
} //main
