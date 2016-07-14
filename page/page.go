// main page structure
package page

import (
	"fmt"
	"math"

	"github.com/AndreaGhizzoni/go-status/config"
	"github.com/shirou/gopsutil/host"
)

type Structure struct {
	Title    string
	UpTime   string
	CPU      Cpu
	RAM      Ram
	DiskStat Disk
	NetStat  Net
}

func formatSeconds(seconds uint64) string {
	minutes := seconds / 60
	s := math.Mod(float64(seconds), 60)
	h := minutes / 60
	m := math.Mod(float64(minutes), 60)
	return fmt.Sprintf("%d:%d:%d", int64(h), int64(m), int64(s))
}

func New(cfg *config.Config) *Structure {
	ut, err := host.Uptime()
	if err != nil {
		panic(err)
	}

	return &Structure{
		Title:    "Rasp status",
		UpTime:   formatSeconds(ut),
		CPU:      *NewCPU(),
		RAM:      *NewRAM(),
		DiskStat: *NewDisk(cfg),
		NetStat:  *NewNet(),
	}
}
