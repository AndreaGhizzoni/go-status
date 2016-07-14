// this file represents the disk part of package page
package page

import (
	"github.com/AndreaGhizzoni/go-status/config"
	"github.com/AndreaGhizzoni/go-status/constant"
	"github.com/AndreaGhizzoni/go-status/util"
	"github.com/shirou/gopsutil/disk"
)

// paths to watch
/*var paths = []string{
	"/home/andrea",
}*/

type PathStat struct {
	Fstype      string
	Total       string
	Free        string
	Used        string
	UsedPercent string
}

type Disk struct {
	Partitions []disk.PartitionStat
	Paths      map[string]PathStat
}

func NewDisk(cfg *config.Config) *Disk {
	partitions, _ := disk.Partitions(false)
	// TODO manage erros

	mappingPaths := make(map[string]PathStat)

	for _, path := range cfg.Watchedpaths {
		stat, err := disk.Usage(path)
		if err == nil {
			mappingPaths[path] = PathStat{
				Fstype:      stat.Fstype,
				Total:       util.Format(stat.Total, constant.Gigabyte),
				Free:        util.Format(stat.Free, constant.Gigabyte),
				Used:        util.Format(stat.Used, constant.Gigabyte),
				UsedPercent: util.FormatPercentege(stat.UsedPercent),
			}
		}
	}

	return &Disk{
		Partitions: partitions,
		Paths:      mappingPaths,
	}
}
