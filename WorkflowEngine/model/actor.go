package model

//Actor - Model for Actor
type Actor struct {
	ActorID string   `json:"actorId" bson:"_id"`
	Role    []string `json:"role" bson:"role"`
	StateID string   `json:"stateId" bson:"stateId"`
}
