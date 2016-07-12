// this file represents the cpu part of the package page
package page

import "github.com/shirou/gopsutil/cpu"

type Core struct {
	CPU        int32
	PhysicalID string
	CoreID     string
	Mhz        float64
}

type Cpu struct {
	Vendor    string
	Model     string
	Family    string
	Cache     int32
	CoreCount int
	Cores     []Core
}

func NewCPU() *Cpu {
	cpuInfo, _ := cpu.Info()
	// TODO check errors

	cores := make([]Core, len(cpuInfo))
	for i, elem := range cpuInfo {
		cores[i] = Core{
			CPU:        elem.CPU,
			PhysicalID: elem.PhysicalID,
			CoreID:     elem.CoreID,
			Mhz:        elem.Mhz,
		}
	}

	cpuCount, _ := cpu.Counts(true)
	// TODO check errors

	return &Cpu{
		Vendor:    cpuInfo[0].VendorID,
		Model:     cpuInfo[0].ModelName,
		Family:    cpuInfo[0].Family,
		Cache:     cpuInfo[0].CacheSize,
		CoreCount: cpuCount,
		Cores:     cores,
	}
}
