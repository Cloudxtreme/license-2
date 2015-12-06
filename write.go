package license

import (
	"errors"
	"io"
)

// Type of license.
type Type int

// The different Licenses.
const (
	MIT = Type(iota)
	BSD = Type(iota)
)

// Write a license with the year and name to the Writer. If there is an error
// writing to the Writer it will be returned.
func Write(lic Type, info *Info, out io.Writer) error {
	// The full license
	var full string

	// Check which license we're writing
	switch lic {
	case MIT:
		full = info.replace(mitText)
	case BSD:
		full = info.replace(bsdText)
	default:
		return errors.New("write license error: invalid type")
	} //switch

	// Write the the Writer
	_, err := out.Write([]byte(full))

	// Return the error
	return err
} //Write
