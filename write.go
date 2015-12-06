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
)

// Write a license with the year and name to the Writer. If there is an error
// writing to the Writer it will be returned.
func Write(lic Type, info *Info, out io.Writer) error {
	// Check which license we're writing
	switch lic {
	case MIT:
		// Write the MIT license
		return writeMIT(info, out)
	} //switch

	return errors.New("write license error: invalid type")
} //Write
