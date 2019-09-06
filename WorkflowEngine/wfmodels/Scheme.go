package wfmodels

import "time"

//WorkflowScheme : Represent Workflow Scheme
type WorkflowScheme struct {
	SchemeID   string            `json:"schemeId"`
	SchemeCode string            `json:"schemeCode"`
	SchemeName string            `json:"schemeName"`
	StartDate  *time.Time        `json:"startDate"`
	EndDate    *time.Time        `json:"endDate"`
	SequenceNo int               `json:"sequenceNo"`
	Actors     []Actor           `json:"actors"`
	Processes  []WorkflowProcess `json:"processes"`
	State      State             `json:"state"`
}
