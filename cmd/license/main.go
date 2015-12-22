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
	// Create the flag to list the available licenses
	list := flag.Bool("list", false,
		"Lists the available licenses instead of writing a license file")

	// Create the flag for the name
	name := flag.String("holder", defaultName, "The name of the copyright holder")

	// Create the flag for the year
	year := flag.Int("year", defaultYear, "The year of the copyright")

	// Create the flag for the license
	lic := flag.String("name", defaultLicense, "The license")

	// Parse the flags
	flag.Parse()

	// Check the list flag
	if *list {
		fmt.Println("Available licenses:\n\tMIT\n\tBSD")
		return
	} //if

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
	license.Write(strings.ToLower(*lic), info, file)
} //main
