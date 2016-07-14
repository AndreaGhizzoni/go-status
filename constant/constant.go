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

	Kilobyte uint64 = 1024
	Megabyte uint64 = 1048576
	Gigabyte uint64 = 1073741824
	Terabyte uint64 = 1099511627776
)
