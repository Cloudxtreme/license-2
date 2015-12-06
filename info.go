package license

import (
	"fmt"
	"strings"
)

// Info for the copyright.
type Info struct {
	Name string
	Year int
} //struct Info

// NewInfo constructs an Info.
func NewInfo(name string, year int) *Info {
	return &Info{Name: name, Year: year}
} //NewInfo

func (info *Info) replace(license string) string {
	// Replace the name
	license = strings.Replace(license, "NAME", info.Name, -1)

	// Replace the year
	return strings.Replace(license, "YEAR", fmt.Sprintf("%d", info.Year), -1)
} //replace
