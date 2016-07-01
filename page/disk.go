// TODO add doc & description
package page

import "github.com/shirou/gopsutil/disk"

// paths to watch
var paths = []string{
	"/home/andrea",
}

type Disk struct {
	Paths      map[string]*disk.UsageStat
	Partitions []disk.PartitionStat
}

func NewDisk() *Disk {
	partitions, _ := disk.Partitions(false)
	// TODO manage erros

	d := &Disk{
		Paths:      make(map[string]*disk.UsageStat),
		Partitions: partitions,
	}

	for _, path := range paths {
		stat, err := disk.Usage(path)
		if err == nil {
			d.Paths[path] = stat
		}
	}

	return d
}
