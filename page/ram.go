// TODO add doc
package page

import "github.com/shirou/gopsutil/mem"

type Ram struct {
	Total       string
	Available   string
	Used        string
	UsedPercent string
	Free        string
}

func NewRAM() *Ram {
	vmstat, _ := mem.VirtualMemory()
	// TODO check err

	return &Ram{
		Total:       Format(vmstat.Total, Megabyte),
		Available:   Format(vmstat.Available, Megabyte),
		Used:        Format(vmstat.Used, Megabyte),
		UsedPercent: FormatPercentege(vmstat.UsedPercent),
		Free:        Format(vmstat.Free, Megabyte),
	}
}
