package constant

import "os/user"

var (
	// application home directory
	AppHome = userDir() + "/.gos/"
	// default application log path
	LogPath = AppHome + "gos.log"
	// defatul configuration file
	ConfigPath = AppHome + "gos.json"
)

const (
	// application server port
	Port = ":8080"

	// template home directory
	TemplateDir = "template/"

	// how many bytes is 1 Kb
	Kilobyte uint64 = 1024
	// how many bytes is 1 Mb
	Megabyte uint64 = 1048576
	// how many bytes is 1 Gb
	Gigabyte uint64 = 1073741824
	// how many bytes is 1 Tb
	Terabyte uint64 = 1099511627776
)

// package function to get current user safe
func userDir() string {
	u, err := user.Current()
	if err != nil {
		panic(err) // this must be a critical error
	}
	return u.HomeDir
}
