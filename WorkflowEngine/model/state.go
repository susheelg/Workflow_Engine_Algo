package model

//State - Model for State
type State struct {
	StateID   string `json:"stateId" bson:"_id"`
	StateName string `json:"stateName" bson:"stateName"`
}
