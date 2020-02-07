package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const SERVER_PORT = ":3000"
//const STATIC_DIR = "/static/"  // for use w/ ReactJS

func main() {
	// Perform all setups before server goes live and starts listening for requests
	initialSetup()
	// Start server listening on 127.0.0.1:3000
	//routeHandler := initRouter()
	fmt.Println("Server started on 127.0.0.1" + SERVER_PORT)
	//startServer(routeHandler)
	startServerReact()
}

func startServerReact() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(SERVER_PORT, nil)
}

func initialSetup() {
	fmt.Println("Initializing services.")
	logger("_", "Service initialized. Executing initialSetup().")
}

// TODO: Find alternative logging implementation that doesn't require opening and closing the log file every time.
func logger(flag string, msg string) {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("[CRITICAL] Error during logger() execution. Stopping program execution.")
		os.Exit(1)
	}
	//defer file.Close()
	log.SetOutput(file)
	//log.Printf(msg)

	switch flag {
	case "w":
		// Warning case
		log.Printf("[WARN]", msg)
	case "e":
		// Error case
		log.Printf("[ERROR]", msg)
	case "c":
		// Critical case
		log.Printf("[CRITICAL]", msg)
	default:
		log.Printf(msg)
	}

	file.Close()
}

// @below is what was used before trying to host w/ ReactJS frontend

//func initRouter() *mux.Router {
//	logger("_", "Executing initRouter().")
//	// Init router
//	r := mux.NewRouter()
//	// Init route handlers
//	r.HandleFunc("/", page_home).Methods("GET")
//	r.HandleFunc("/", page_about).Methods("GET")
//	http.Handle("/", r)
//	return r
//}

//func startServer(r *mux.Router) {
//	fmt.Println("Initializing web server.")
//	logger("_", "Executing startServer().")
//	http.ListenAndServe(SERVER_PORT, r)
//
//	// Server CSS, JS & Images Statically (for use w/ ReactJS)
//	//router.
//	//	PathPrefix(STATIC_DIR).
//	//	Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))
//}

//func page_home(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "<h1>Hey, a H1 header!</h1>")
//}
//
//func page_about(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "<h4>Welcome to the ABOUT page!</h4>")
//}
