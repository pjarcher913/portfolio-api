/*
AUTHOR: Patrick Archer (@pjarcher913)
DATE CREATED: 5 May 2020
Copyright (c) 2020 Patrick Archer
*/

/*
TODO: Connect to database one time and defer until program close instead of having to connect during each PRH.
	Can't just call db.Open() from a separate method because upon method return the db closes.
 */

package api

import (
	"../../models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Database params
const (
	host     = "localhost"
	port     = 5432
	user     = "guest"
	password = "guest"
	dbname   = "portfolio-api"
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

	/*
	Connect to database and query data to be returned in future HTTP response.
	 */

	// Connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Infoln("Database connection to " + dbname + " initialized successfully.")

	// Database query statement
	query := `SELECT msg FROM "default".api WHERE api_id = 1;`

	// Query data to be returned to user
	var msg string
	err = db.QueryRow(query).Scan(&msg)
	if err != nil {
		log.Fatalln(err)
	}

	/*
	Prepare HTTP response
	 */

	w.Header().Set("Content-Type", "application/json")

	// Get raw request URL path
	reqUrl := r.URL

	// Get request params
	params := mux.Vars(r)

	// Populate response struct
	response := models.Model_POST_1{
		Message:   msg,
		Parameter: params["rootParam"],
		Timestamp: time.Now().UTC().String(),
	}

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
