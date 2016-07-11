package page

import "fmt"

const (
	Kilobyte uint64 = 1024
	Megabyte uint64 = 1048576
	Gigabyte uint64 = 1073741824
	Terabyte uint64 = 1099511627776
)

// i must be represented in bytes
func Format(i uint64, level uint64) string {
	if i >= Terabyte && level == Terabyte {
		return fmt.Sprintf("%d Tb", i/Terabyte)
	} else if i >= Gigabyte && level == Gigabyte {
		return fmt.Sprintf("%d Gb", i/Gigabyte)
	} else if i >= Megabyte && level == Megabyte {
		return fmt.Sprintf("%d Mb", i/Megabyte)
	} else if i >= Kilobyte && level == Kilobyte {
		return fmt.Sprintf("%d Kb", i/Kilobyte)
	} else {
		return "error form format"
	}
}
