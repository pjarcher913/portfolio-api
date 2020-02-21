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
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"models"
	"net/http"
	"os"
	"time"
)

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
	r.NotFoundHandler = http.HandlerFunc(prh_404)
	/* GETs */
	r.HandleFunc("/", prh_GET_Home).Methods("GET")
	r.HandleFunc("/resume", prh_GET_Resume).Methods("GET")
	r.HandleFunc("/projects", prh_GET_Projects).Methods("GET")
	r.HandleFunc("/contact", prh_GET_Contact).Methods("GET")
	/* POSTs */
	r.HandleFunc("/{rootParam}", prh_POST_Home).Methods("POST")

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

// TODO: Can't serve after w.WriteHeader(), so figure out a way to send 404 code and custom 404.html
//prh_404() is the website's "404 Error" handler when users try to navigate to an invalid/un-served route.
func prh_404(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing prh_404().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.Error(w, "Error(404): Page not found!", 404)	// try to find a way to not use this
	//w.WriteHeader(http.StatusNotFound)
	//fmt.Println(w.Header().Get("status"))
	//w.Header().Set("status", "404 Not Found")
	//http.ServeFile(w, r, PATH_TO_404_HTML)
}

// prh_GET_Home() is the website's "Home" page GET route handler.
func prh_GET_Home(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing prh_GET_Home().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_HOME_HTML)
}

// prh_GET_Resume() is the website's "Resume" page GET route handler.
func prh_GET_Resume(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing prh_GET_Resume().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_RESUME_HTML)
}

// prh_GET_Projects() is the website's "Projects" page GET route handler.
func prh_GET_Projects(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing prh_GET_Projects().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_PROJECTS_HTML)
}

// prh_GET_Contact() is the website's "Contact" page GET route handler.
func prh_GET_Contact(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing prh_GET_Contact().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_CONTACT_HTML)
}

// prh_POST_Home() is the website's "Home" page POST route handler.
// This is an Easter Egg!
func prh_POST_Home(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing prh_POST_Home(), which is an Easter Egg!")
	w.Header().Set("Content-Type", "application/json")

	// Get raw request URL path
	reqUrl := r.URL

	// Get request params
	params := mux.Vars(r)

	// Populate response struct
	response := models.Model_EasterEgg{
		Message:   "Hey, you found an API Easter Egg!",
		Parameter: params["rootParam"],
		Timestamp: time.Now().UTC().String(),
	}

	// Log response
	log.WithFields(log.Fields{
		"responseData": response,
		"allParams": params,
		"fullURL": reqUrl,
	}).Debug("RESPONSE-prh_POST_Home()")

	// Encode response as JSON and send to client via http.ResponseWriter
	encodingErr := json.NewEncoder(w).Encode(response)
	if encodingErr != nil {
		log.Fatalln(encodingErr.Error())
	}
}
