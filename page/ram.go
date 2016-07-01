// TODO add doc
package page

import "github.com/shirou/gopsutil/mem"

type Ram struct {
	Total       uint64
	Available   uint64
	Used        uint64
	UsedPercent float64
	Free        uint64
}

func NewRAM() *Ram {
	vmstat, _ := mem.VirtualMemory()
	// TODO check err

	return &Ram{
		Total:       vmstat.Total,
		Available:   vmstat.Available,
		Used:        vmstat.Used,
		UsedPercent: vmstat.UsedPercent,
		Free:        vmstat.Free,
	}
}
