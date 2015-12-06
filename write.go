package license

import (
	"errors"
	"io"
)

// Write a license with the year and name to the Writer. If there is an error
// writing to the Writer it will be returned.
func Write(license string, info *Info, out io.Writer) error {
	// The full license
	var full string

	// Check which license we're writing
	switch license {
	case "mit":
		full = info.replace(mitText)
	case "bsd":
		full = info.replace(bsdText)
	default:
		return errors.New("write license error: invalid type")
	} //switch

	// Write the the Writer
	_, err := out.Write([]byte(full))

	// Return the error
	return err
} //Write
