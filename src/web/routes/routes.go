/*
AUTHOR: Patrick Archer (@pjarcher913)
DATE CREATED: 5 May 2020
Copyright (c) 2020 Patrick Archer
*/

package routes

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

/*======================================================================================*/
// Declarations

const BUILD_PATH = "src/web/pages/home"  // Location of frontend UI React build to be served
const PATH_TO_404_HTML = "../src/web/pages/errors/404.html"  // Location of 404.html (used to render page)

/*======================================================================================*/
// 404 Error Handler

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

/*======================================================================================*/
// GET Handlers

// TODO

// [!] DEPRECATED, keeping for future reference [!]
//PRH_GET_Home() is a GET route handler for the path $HOST$/api/1.
//func PRH_GET_Home(w http.ResponseWriter, r *http.Request) {
//	log.Infoln("Executing PRH_GET_1().")
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//
//	//buildHandler := http.FileServer(http.Dir(BUILD_PATH))
//	////http.Handle("/", buildHandler)
//	//http.ListenAndServe("/", buildHandler)
//
//	//http.ServeContent(w, r, http.FileServer(http.Dir(BUILD_PATH)))
//	//http.ServeFile(w, r, http.FileServer(http.Dir(BUILD_PATH))
//
//	//http.ServeFile(w, r, PATH_TO_HOME_HTML)
//}
