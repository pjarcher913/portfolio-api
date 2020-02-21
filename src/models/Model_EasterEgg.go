/*
AUTHOR: Patrick Archer (@pjarcher913)
DATE CREATED: 7 February 2020
Copyright (c) Patrick Archer as of 7 February 2020
*/

package models

import ()

// Model_EasterEgg defines a data struct that will be used in responses to applicable API requests
type Model_EasterEgg struct {
	// Define properties of struct
	Message 	string	`json:"msg"`
	Parameter 	string	`json:"param"`
	Timestamp 	string	`json:"time"`
}
