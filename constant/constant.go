package constant

const (
	// application server port
	Port = ":8080"

	// application home directory
	AppHome = "/home/andrea/.gos/"
	// default application log path
	LogPath = AppHome + "gos.log"
	// defatul configuration file
	ConfigPath = AppHome + "gos.json"
	// template home directory
	TemplateDir = AppHome + "template/"

	// how many bytes is 1 Kb
	Kilobyte uint64 = 1024
	// how many bytes is 1 Mb
	Megabyte uint64 = 1048576
	// how many bytes is 1 Gb
	Gigabyte uint64 = 1073741824
	// how many bytes is 1 Tb
	Terabyte uint64 = 1099511627776
)
