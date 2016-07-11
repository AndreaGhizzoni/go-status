// TODO add doc & description
package page

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

// paths to watch
var paths = []string{
	"/home/andrea",
}

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

func NewDisk() *Disk {
	partitions, _ := disk.Partitions(false)
	// TODO manage erros

	mappingPaths := make(map[string]PathStat)

	for _, path := range paths {
		stat, err := disk.Usage(path)
		if err == nil {
			mappingPaths[path] = PathStat{
				Fstype:      stat.Fstype,
				Total:       Format(stat.Total, Gigabyte),
				Free:        Format(stat.Free, Gigabyte),
				Used:        Format(stat.Used, Gigabyte),
				UsedPercent: fmt.Sprintf("%.2f", stat.UsedPercent),
			}
		}
	}

	return &Disk{
		Partitions: partitions,
		Paths:      mappingPaths,
	}
}
