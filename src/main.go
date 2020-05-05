/*
AUTHOR: Patrick Archer (@pjarcher913)
DATE CREATED: 7 February 2020
Copyright (c) 2020 Patrick Archer
*/

/*
TODO: Maintain README.md
TODO: Create more API EEs
TODO: Prevent users from spamming API calls and crashing the program (keep efficiency in mind when drafting a solution).
TODO: Create .bat to create `~/logs` (if it doesn't exist) and execute `BUILD.exe`. Can include other functionalities as needed.
*/

package main

import (
	"../src/web/api"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

/*======================================================================================*/
// Declarations

// Only enable if debugging and want extra log info recorded
const DEBUG_MODE = true  // Limits functionality for dev. purposes

const LOG_PATH = "./logs/"  // Where log files are to be generated and stored

const SERVER_PORT = ":3000"  // Host port to serve on
const PATH_TO_HOME_HTML = "./src/web/pages/home/home.html"  // Location of home.html (used to render page)
const PATH_TO_RESUME_HTML = "./src/web/pages/resume/resume.html"  // Location of resume.html (used to render page)
const PATH_TO_PROJECTS_HTML = "./src/web/pages/projects/projects.html"  // Location of projects.html (used to render page)
const PATH_TO_CONTACT_HTML = "./src/web/pages/contact/contact.html"  // Location of contact.html (used to render page)
const PATH_TO_404_HTML = "./src/web/pages/errors/404.html"  // Location of 404.html (used to render page)

var LOG_STAMP = "main_" + xid.New().String()  // Unique id tag included into newly-generated log file names

/*======================================================================================*/
// Main

func main() {
	// Perform all preliminary setups before server goes live and starts listening for requests.
	// If prelimSetup() completes, it returns true and the program continues.
	// If it fails, it returns false and fatal error.
	if prelimSetup() {
		// Initialize route handler
		routeHandler := initRouter()
		// Begin serving the site and listening on assigned web_old port
		initWebServer(routeHandler)
	} else {
		log.Fatalln("prelimSetup() failed.")
	}
}

/*======================================================================================*/
// Additional Functions

// initLogger() initializes Logrus and configures it for future utilization.
func initLogger() {
	file, err := os.OpenFile(LOG_PATH + LOG_STAMP + ".log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//defer file.Close()  // Because file will be closed in GC, we can leave it open so the logger is uninterrupted
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})

	// Minimum severity level to log, according to if the program is running in debug mode or not
	if DEBUG_MODE { log.SetLevel(log.DebugLevel) } else { log.SetLevel(log.ErrorLevel) }

	log.Infoln("Logger initialized successfully.")
}

// prelimSetup() performs all preliminary setups before server goes live and starts listening for requests.
func prelimSetup() bool {
	// Initialize logger
	initLogger()

	// Services init'd, thus prelimSetup() is complete
	log.Println("Initializing services via prelimSetup().")
	return true
}

// initRouter() initializes Mux's routing services and configures them according to the website's defined route hierarchy.
func initRouter() *mux.Router {
	log.Infoln("Executing initRouter().")

	// Init mux router object
	r := mux.NewRouter()

	// Init route handlers
	/* 404 */
	r.NotFoundHandler = http.HandlerFunc(api.PRH_404)
	/* GETs */
	// TODO: create directory and structure for normal page route handlers and separate from backend API handlers
	r.HandleFunc("/api/1", api.PRH_GET_1).Methods("GET")
	r.HandleFunc("/api/2", api.PRH_GET_2).Methods("GET")
	r.HandleFunc("/api/3", api.PRH_GET_3).Methods("GET")
	r.HandleFunc("/api/4", api.PRH_GET_4).Methods("GET")
	/* POSTs */
	//r.HandleFunc("/{rootParam}", prh_POST_Home).Methods("POST")
	r.HandleFunc("/api/1/{rootParam}", api.PRH_POST_1).Methods("POST")

	return r
}

// initWebServer() initializes the web server and begins serving clients connecting to the pre-configured SERVER_PORT
func initWebServer(routeHandler *mux.Router) {
	log.Infoln("Executing initWebServer().")

	// Serve website and listen on configured SERVER_PORT
	// http.ListenAndServe() always returns a non-nil error, and the error is its only returned value.
	// However, http.ListenAndServe() should never return (unless error or intentional termination).
	fmt.Println("Now serving on 127.0.0.1" + SERVER_PORT)
	log.Infoln("Now serving on 127.0.0.1" + SERVER_PORT)
	err := http.ListenAndServe(SERVER_PORT, routeHandler)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

