/*
AUTHOR: Patrick Archer (@pjarcher913)
DATE CREATED: 7 February 2020
Copyright (c) 2020 Patrick Archer
*/

/*
TODO: Maintain README.md and API_DOCUMENTATION.md.
TODO: Update LICENSE and NOTICE.md to reflect new developments.
TODO: Create more API EEs.
TODO: Prevent users from spamming API calls and crashing the program (keep efficiency in mind when drafting a solution).
TODO: Create .bat to create `~/logs` (if it doesn't exist) and execute `BUILD.exe`. Can include other functionalities as needed.
*/

package main

import (
	"./web/api"
	"./web/routes"
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
const BUILD_PATH = "src/web/pages/home"  // Location of frontend UI React build to be served

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

	// Init route handlers -- MUST init frontend catch-all "/" handler LAST

	/* 404 */
	r.NotFoundHandler = http.HandlerFunc(routes.PRH_404)

	/* API */
	apiRoute := r.PathPrefix("/api/").Subrouter()
	apiRoute.HandleFunc("/1/{rootParam}", api.PRH_POST_1).Methods("POST")

	/*  FRONTEND & SUPPORTING FILES -- MUST INIT LAST */
	buildHandler := http.FileServer(http.Dir(BUILD_PATH))
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("/static")))
	r.PathPrefix("/").Handler(buildHandler)
	r.PathPrefix("/static/").Handler(staticHandler)

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

