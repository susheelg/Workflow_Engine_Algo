package model

import "time"

//Process - Model for a Process
type Process struct {
	ProcessID       string    `json:"processID" bson:"_id"`
	ProcessName     string    `json:"processName" bson:"processName"`
	ProcessCode     string    `json:"processCode" bson:"processCode"`
	SchemeID        string    `json:"schemeID" bson:"schemeID"`
	StartDate       time.Time `json:"startDate" bson:"startDate"`
	EndDate         time.Time `json:"endDate" bson:"endDate"`
	StateID         string    `json:"stateID" bson:"stateID"`
	Actors          []Actor   `json:"actors" bson:"actors"`
	PreviousProcess string    `json:"previousProcess" bson:"previousProcess"`
	NextProcess     string    `json:"nextProcess" bson:"nextProcess"`
	RootProcess     string    `json:"rootProcess" bson:"rootProcess"`
	RootActivityID  string    `json:"rootActivityID" bson:"rootActivityID"`
	IsAsyncProcess  bool      `json:"isAsyncProcess" bson:"isAsyncProcess"`

	ProcessInvokedByActivityID string `json:"processInvokedByActivityID" bson:"processInvokedByActivityID"`
}
