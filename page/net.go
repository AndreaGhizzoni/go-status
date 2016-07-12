// this file represents the net part of package page
package page

import "github.com/shirou/gopsutil/net"

type Net struct {
	Interfaces []net.InterfaceStat
}

func NewNet() *Net {
	inter, _ := net.Interfaces()
	// TODO check errors

	return &Net{
		Interfaces: inter,
	}
}
