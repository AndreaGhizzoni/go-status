// TODO add doc
package page

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type Ram struct {
	Total       string
	Available   string
	Used        string
	UsedPercent string
	Free        string
}

const (
	kilobyte uint64 = 1024
	megabyte uint64 = 1048576
	gigabyte uint64 = 1073741824
)

// i must be represented in bytes
func format(i uint64) string {
	/*if i >= gigabyte {
		return fmt.Sprintf("%d Gb", i/gigabyte)
	} else
	*/
	if i >= megabyte {
		return fmt.Sprintf("%d Mb", i/megabyte)
	} else if i >= kilobyte {
		return fmt.Sprintf("%d Kb", i/kilobyte)
	} else {
		return "error"
	}
}

func NewRAM() *Ram {
	vmstat, _ := mem.VirtualMemory()
	// TODO check err

	return &Ram{
		Total:       format(vmstat.Total),
		Available:   format(vmstat.Available),
		Used:        format(vmstat.Used),
		UsedPercent: fmt.Sprintf("%.2f", vmstat.UsedPercent),
		Free:        format(vmstat.Free),
	}
}
