package license

// Info for the copyright.
type Info struct {
	Name string
	Year int
} //struct Info

// NewInfo constructs an Info.
func NewInfo(name string, year int) *Info {
	return &Info{Name: name, Year: year}
} //NewInfo
