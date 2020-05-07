/*
AUTHOR: Patrick Archer (@pjarcher913)
DATE CREATED: 7 February 2020
Copyright (c) Patrick Archer as of 7 February 2020
*/

package models

import ()

// Model_EasterEgg defines a data struct that will be used in responses to PRH_POST_1 API requests
type Model_POST_1 struct {
	// Custom text message
	Message 	string	`json:"msg"`
	// Parameter in request URL
	Parameter 	string	`json:"param"`
	// Timestamp when response is created
	Timestamp 	string	`json:"time"`
}
