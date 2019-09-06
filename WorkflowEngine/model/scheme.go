package model

import "time"

//Scheme - Model for a scheme
type Scheme struct {
	SchemeID        string    `json:"schemeId" bson:"_id"`
	SchemeCode      string    `json:"schemeCode" bson:"schemeCode"`
	SchemeName      string    `json:"schemeName" bson:"schemeName"`
	SchemeStartDate time.Time `json:"schemeStartDate" bson:"schemeStartDate"`
	SchemeEndDate   time.Time `json:"schemeEndDate" bson:"schemeEndDate"`
	Actors          []Actor   `json:"actors" bson:"actors"`
	Processes       []string  `json:"processes" bson:"processes"`
	SequenceNo      int       `json:"sequenceNo" bson:"sequenceNo"`
	StateID         string    `json:"stateId" bson:"stateId"`
}
