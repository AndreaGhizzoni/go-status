//TODO add doc & description
package page

import (
	"math"

	"fmt"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
)

type Structure struct {
	Title          string
	UpTime         string
	CPU            Cpu
	RAM            Ram
	DiskUsagePaths map[string]*disk.UsageStat
	Partitions     []disk.PartitionStat
	Interfaces     []net.InterfaceStat
}

func formatSeconds(seconds uint64) string {
	minutes := seconds / 60
	s := math.Mod(float64(seconds), 60)
	h := minutes / 60
	m := math.Mod(float64(minutes), 60)
	return fmt.Sprintf("%d:%d:%d", int64(h), int64(m), int64(s))
}

func panicIf(e error) {
	if e != nil {
		panic(e)
	}
}

// TODO refactoring this method into NewPageStructure
func New() *Structure {
	ut, err := host.Uptime()
	panicIf(err)

	diskPaths := make(map[string]*disk.UsageStat)
	path := "/home/andrea"
	stat, err := disk.Usage(path)
	if err == nil {
		diskPaths[path] = stat
	}

	p, err := disk.Partitions(false)
	panicIf(err)

	interf, err := net.Interfaces()
	panicIf(err)

	return &Structure{
		Title:          "Rasp status",
		UpTime:         formatSeconds(ut),
		CPU:            *NewCPU(),
		RAM:            *NewRAM(),
		DiskUsagePaths: diskPaths,
		Partitions:     p,
		Interfaces:     interf,
	}
}
