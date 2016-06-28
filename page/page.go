//TODO add doc
package page

import (
	"math"

	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Structure struct {
	Title          string
	UpTime         string
	CPUCount       int
	CPUModel       string
	CPUCache       int32
	CPUsStat       []cpu.InfoStat
	CPULoad        *load.AvgStat
	VMStat         *mem.VirtualMemoryStat
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

func New() *Structure {
	ut, _ := host.Uptime()

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

	interf, _ := net.Interfaces()

	return &Structure{
		Title:          "Rasp status",
		UpTime:         formatSeconds(ut),
		CPUCount:       cpuCounts,
		CPUModel:       cpuInfo[0].ModelName,
		CPUCache:       cpuInfo[0].CacheSize,
		CPUsStat:       cpuInfo,
		CPULoad:        l,
		VMStat:         vmstat,
		DiskUsagePaths: diskPaths,
		Partitions:     p,
		Interfaces:     interf,
	}
}
