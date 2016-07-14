// this file represents the ram part of the package page
package page

import (
	"github.com/AndreaGhizzoni/go-status/constant"
	"github.com/AndreaGhizzoni/go-status/util"
	"github.com/shirou/gopsutil/mem"
)

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
		Total:       util.Format(vmstat.Total, constant.Megabyte),
		Available:   util.Format(vmstat.Available, constant.Megabyte),
		Used:        util.Format(vmstat.Used, constant.Megabyte),
		UsedPercent: util.FormatPercentege(vmstat.UsedPercent),
		Free:        util.Format(vmstat.Free, constant.Megabyte),
	}
}
