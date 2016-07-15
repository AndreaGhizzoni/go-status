// main package that contains the web server
package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/AndreaGhizzoni/go-status/config"
	"github.com/AndreaGhizzoni/go-status/constant"
	"github.com/AndreaGhizzoni/go-status/page"
	"github.com/AndreaGhizzoni/go-status/util"
)

var (
	showVersion = false
	version     = "0.1.1"
)

var (
	TraceF *log.Logger
	ErrorF *log.Logger
	ErrorC *log.Logger
)

var (
	// Variable to hold all html template files
	html = template.Must(template.ParseGlob(constant.TemplateDir + "*.html"))

	// Variable to hold configurator pointer for all web handler
	cfg *config.Config
)

// function to initialize the logger
func initLogger() {
	flag := log.Ldate | log.Ltime | log.Lshortfile
	ErrorC = log.New(os.Stderr, "ERROR: ", flag)

	logFile, err := os.OpenFile(constant.LogPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		ErrorC.Fatalln("Failed to open log file :", err)
	}

	TraceF = log.New(logFile, "TRACE: ", flag)
	ErrorF = log.New(logFile, "ERROR: ", flag)
}

// function to check if home directory exists, if not create a new one
func initHomeDir() {
	if e, _ := util.ExistsPath(constant.AppHome); !e {
		err := os.Mkdir(constant.AppHome, 0700)
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
	TraceF.Print("Index handler called")
	renderTemplate(w, "index", page.New(cfg))
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	TraceF.Print("Shutdown handler called")
	err := exec.Command("/sbin/poweroff").Run()
	if err != nil {
		ErrorF.Fatal(err)
	}
}

func main() {
	flag.BoolVar(&showVersion, "version", false, "show the current version")
	flag.Parse()

	if showVersion {
		fmt.Println(version)
	} else {
		initHomeDir()
		initLogger()

		c, err := config.Parse()
		if err != nil {
			ErrorF.Fatal(err)
		}
		cfg = c

		TraceF.Print("Program Start")
		http.HandleFunc("/", rootHandler)
		http.HandleFunc("/shutdown", shutdownHandler)
		fs := http.FileServer(http.Dir(constant.TemplateDir))
		http.Handle("/template/", http.StripPrefix("/template/", fs))
		TraceF.Print("Server Started")
		ErrorF.Fatal(http.ListenAndServe(constant.Port, nil))
	}
}
