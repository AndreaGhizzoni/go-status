// TODO add package doc
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/AndreaGhizzoni/go-status/page"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

var (
	// application home directory
	appHome = "/home/andrea/.gos/"

	// template home directory
	templateDir = appHome + "template/"

	// Variable to hold all html template files
	html = template.Must(template.ParseGlob(templateDir + "*.html"))

	// default application log path
	defLogPath = appHome + "gos.log"
)

// function to initialize the logger
func initLogger(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	flag := log.Ldate | log.Ltime | log.Lshortfile

	Trace = log.New(traceHandle, "TRACE: ", flag)
	Info = log.New(infoHandle, "INFO: ", flag)
	Warning = log.New(warningHandle, "WARNING: ", flag)
	Error = log.New(errorHandle, "ERROR: ", flag)
}

// function to check if home directory exists, if not create a new one
func initHomeDir() {
	if e, _ := exists(appHome); !e {
		err := os.Mkdir(appHome, 0700)
		if err != nil {
			panic(err)
		}
	}
}

// Auxiliary function to render a specific template.
// w writer where to write the template
// tmpl template name
// p page structure where to find the data
func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Structure) {
	err := html.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Http handler for index
func rootHandler(w http.ResponseWriter, r *http.Request) {
	Trace.Print("Index handler called")
	renderTemplate(w, "index", page.New())
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	Trace.Print("Shutdown handler called")
	err := exec.Command("/sbin/poweroff").Run()
	if err != nil {
		Error.Fatal(err)
	}
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func main() {
	initHomeDir()

	logFile, err := os.OpenFile(defLogPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalln("Failed to open log file :", err)
	}
	defer func() {
		Trace.Print("Closing log file")
		logFile.Close()
	}()

	initLogger(logFile, logFile, logFile, logFile)

	Trace.Print("Program Start")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/shutdown", shutdownHandler)
	fs := http.FileServer(http.Dir(templateDir))
	http.Handle("/template/", http.StripPrefix("/template/", fs))
	Trace.Print("Server Started")
	Error.Fatal(
		http.ListenAndServe(":8080", nil))
}
