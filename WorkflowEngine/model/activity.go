package model

import "time"

//Activity - Model for Activity
type Activity struct {
	ActivityID         string      `json:"activityID" bson:"_id"`
	ActivityCode       string      `json:"activityCode" bson:"activityCode"`
	ProcessID          string      `json:"processID" bson:"processID"`
	ActivityName       string      `json:"activityName" bson:"activityName"`
	StartDate          time.Time   `json:"startDate" bson:"startDate"`
	EndDate            time.Time   `json:"endDate" bson:"endDate"`
	Actors             []Actor     `json:"actors" bson:"actors"`
	PreviousActivityID string      `json:"previousActivityID" bson:"previousActivityID"`
	AcceptActivityID   string      `json:"acceptActivityID" bson:"acceptActivityID"`
	AejectActivityID   string      `json:"rejectActivityID" bson:"rejectActivityID"`
	StateID            string      `json:"stateID" bson:"stateID"`
	ExecutionTime      int64       `json:"executionTime" bson:"executionTime"`
	ActivityInvokedBy  string      `json:"activityInvokedBy" bson:"activityInvokedBy"`
	Executor           interface{} `json:"executor" bson:"executor"`
	IsAsyncActivity    bool        `json:"isAsyncActivity" bson:"isAsyncActivity"`
	SequenceNo         int         `json:"sequenceNo" bson:"sequenceNo"`
}
