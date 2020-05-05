/*
AUTHOR: Patrick Archer (@pjarcher913)
DATE CREATED: 5 May 2020
Copyright (c) 2020 Patrick Archer
*/

package api

import (
	"../../models"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

/*======================================================================================*/
// Declarations

const PATH_TO_404_HTML = "../src/web/pages/errors/404.html"  // Location of 404.html (used to render page)
const PATH_TO_HOME_HTML = "./src/web/pages/home/home.html"  // Location of home.html (used to render page)
const PATH_TO_RESUME_HTML = "./src/web/pages/resume/resume.html"  // Location of resume.html (used to render page)
const PATH_TO_PROJECTS_HTML = "./src/web/pages/projects/projects.html"  // Location of projects.html (used to render page)
const PATH_TO_CONTACT_HTML = "./src/web/pages/contact/contact.html"  // Location of contact.html (used to render page)

/*======================================================================================*/
// GETs

// TODO: Can't serve after w.WriteHeader(), so figure out a way to send 404 code and custom 404.html
//PRH_404() is a "404 Error" route handler when users try to navigate to an invalid/un-served route.
func PRH_404(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing PRH_404().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.Error(w, "Error(404): Page not found!", 404)	// try to find a way to not use this
	//w.WriteHeader(http.StatusNotFound)
	//fmt.Println(w.Header().Get("status"))
	//w.Header().Set("status", "404 Not Found")
	//http.ServeFile(w, r, PATH_TO_404_HTML)
}

// PRH_GET_1() is a GET route handler for the path $HOST$/api/1.
func PRH_GET_1(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing PRH_GET_1().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_HOME_HTML)
}

// PRH_GET_2() is a GET route handler for the path $HOST$/api/2.
func PRH_GET_2(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing PRH_GET_2().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_RESUME_HTML)
}

// PRH_GET_3() is a GET route handler for the path $HOST$/api/3.
func PRH_GET_3(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing PRH_GET_3().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_PROJECTS_HTML)
}

// PRH_GET_4() is a GET route handler for the path $HOST$/api/4.
func PRH_GET_4(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing PRH_GET_4().")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, PATH_TO_CONTACT_HTML)
}

/*======================================================================================*/
// POSTs

// PRH_POST_1() is a POST route handler for the path $HOST$/api/1/{rootParam}.
// This is an Easter Egg!
func PRH_POST_1(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Executing PRH_POST_1().")
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
	}).Debug("RESPONSE-PRH_POST_1()")

	// Encode response as JSON and send to client via http.ResponseWriter
	encodingErr := json.NewEncoder(w).Encode(response)
	if encodingErr != nil {
		log.Fatalln(encodingErr.Error())
	}
}
