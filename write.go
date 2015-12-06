package license

import "io"

// Type of license.
type Type int

// The different Licenses.
const (
	MIT = Type(iota)
)

// Write a license with the year and name to the Writer. If there is an error
// writing to the Writer it will be returned.
func Write(lic Type, info *Info, out io.Writer) error {
	return nil
} //Write
