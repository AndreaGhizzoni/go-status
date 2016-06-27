//TODO add doc
package page

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

type Structure struct {
	Title          string
	CPUCount       int
	CPUModel       string
	CPUCache       int32
	CPUsStat       []cpu.InfoStat
	CPULoad        *load.AvgStat
	VMStat         *mem.VirtualMemoryStat
	DiskUsagePaths map[string]*disk.UsageStat
	Partitions     []disk.PartitionStat
}

func New() *Structure {
	cpuInfo, err := cpu.Info()

	cpuCounts, err := cpu.Counts(true)
	if err != nil {
		cpuCounts = -1
	}

	vmstat, _ := mem.VirtualMemory()

	diskPaths := make(map[string]*disk.UsageStat)
	path := "/home/andrea"
	stat, err := disk.Usage(path)
	if err == nil {
		diskPaths[path] = stat
	}

	p, _ := disk.Partitions(false)

	l, _ := load.Avg()

	return &Structure{
		Title:          "Rasp status",
		CPUCount:       cpuCounts,
		CPUModel:       cpuInfo[0].ModelName,
		CPUCache:       cpuInfo[0].CacheSize,
		CPUsStat:       cpuInfo,
		CPULoad:        l,
		VMStat:         vmstat,
		DiskUsagePaths: diskPaths,
		Partitions:     p,
	}
}
