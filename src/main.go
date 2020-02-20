/*
TODO: Maintain README.md
TODO: Update/verify license
TODO: Create more API EEs
TODO: Prevent users from spamming API calls and crashing the program (keep efficiency in mind when drafting a solution).
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
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
//const PATH_TO_404_HTML = "./src/web/pages/errors/404.html"  // Location of 404.html (used to render page)

var LOG_STAMP = "main_" + xid.New().String()  // Unique id tag included into newly-generated log file names

// TODO: Move to ./src/models and re-implement here in main.go
// ee is an Easter Egg data struct that will be used in response to applicable API requests
type ee struct {
	// Define properties of struct
	Message 	string		`json:"msg"`
	Parameter 	string		`json:"param"`
	Timestamp 	string		`json:"time"`
}

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
	/* GETs */
	r.HandleFunc("/", prh_Home_GET).Methods("GET")
	r.HandleFunc("/resume", prh_Resume_GET).Methods("GET")
	r.HandleFunc("/projects", prh_Projects_GET).Methods("GET")
	r.HandleFunc("/contact", prh_Contact_GET).Methods("GET")
	/* POSTs */
	r.HandleFunc("/{rootParam}", prh_Home_POST).Methods("POST")
	/* 404 */
	//r.NotFoundHandler = custom404Handler

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

// TODO
// custom404Handler() is the website's "404 Error" handler when users try to navigate to an invalid route.
//func custom404Handler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	http.ServeFile(w, r, PATH_TO_404_HTML)
//}

// prh_Home_GET() is the website's "Home" page GET route handler.
func prh_Home_GET(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing prh_Home_GET().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_HOME_HTML)
}

// prh_Resume_GET() is the website's "Resume" page GET route handler.
func prh_Resume_GET(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing prh_Resume_GET().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_RESUME_HTML)
}

// prh_Projects_GET() is the website's "Projects" page GET route handler.
func prh_Projects_GET(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing prh_Projects_GET().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_PROJECTS_HTML)
}

// prh_Contact_GET() is the website's "Contact" page GET route handler.
func prh_Contact_GET(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing prh_Contact_GET().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_CONTACT_HTML)
}

// prh_Home_POST() is the website's "Home" page POST route handler.
// This is an Easter Egg!
func prh_Home_POST(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing prh_Home_POST(), which is an Easter Egg!")
	w.Header().Set("Content-Type", "application/json")

	// Get raw request URL path
	reqUrl := r.URL

	// Get request params
	params := mux.Vars(r)

	// Populate response struct
	easterEgg := ee {
		Message: "Hey, you found an Easter Egg!",
		Parameter: params["rootParam"],
		Timestamp: time.Now().UTC().String(),
	}

	// Log response
	log.WithFields(log.Fields{
		"responseData": easterEgg,
		"params": params,
		"url": reqUrl,
	}).Debug("RESPONSE-prh_Home_POST()")

	// Encode response as JSON and send to client
	json.NewEncoder(w).Encode(easterEgg)
}
