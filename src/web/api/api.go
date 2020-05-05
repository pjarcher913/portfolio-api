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
// GETs

// TODO

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
