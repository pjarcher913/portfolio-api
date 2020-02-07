package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const SERVER_PORT = ":3000"  // Host port to serve on
const WEB_DIR = "./src/web"  // Location of the static production build directory for the website frontend

func main() {
	// Perform all setups before server goes live and starts listening for requests
	initialSetup()
	// Begin serving the site and listening on assigned web port
	fmt.Println("Server started on 127.0.0.1" + SERVER_PORT)
	startServer()
}

func startServer() {
	http.Handle("/", http.FileServer(http.Dir(WEB_DIR)))
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
		fmt.Printf("Terminating program. Error during logger() execution:\n" + err.Error())
		os.Exit(1)
	}
	log.SetOutput(file)

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
