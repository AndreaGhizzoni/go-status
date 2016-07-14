// TODO add doc
package util

import (
	"fmt"
	"os"

	"github.com/AndreaGhizzoni/go-status/constant"
)

// exists returns whether the given file or directory exists or not
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// i must be represented in bytes
func Format(i uint64, level uint64) string {
	if i >= constant.Terabyte && level == constant.Terabyte {
		return fmt.Sprintf("%d Tb", i/constant.Terabyte)
	} else if i >= constant.Gigabyte && level == constant.Gigabyte {
		return fmt.Sprintf("%d Gb", i/constant.Gigabyte)
	} else if i >= constant.Megabyte && level == constant.Megabyte {
		return fmt.Sprintf("%d Mb", i/constant.Megabyte)
	} else if i >= constant.Kilobyte && level == constant.Kilobyte {
		return fmt.Sprintf("%d Kb", i/constant.Kilobyte)
	} else {
		return "error form format"
	}
}

func FormatPercentege(p float64) string {
	return fmt.Sprintf("%.2f", p)
}
