package wfmodels

import "time"

//WorkflowActivity : Represent Workflow Activity
type WorkflowActivity struct {
	ActivityID         string          `json:"activityId"`
	ActivityCode       string          `json:"activityCode"`
	ActivityName       string          `json:"activityName"`
	WorkflowScheme     WorkflowScheme  `json:"workflowScheme"`
	WorkflowProcess    WorkflowProcess `json:"workflowProcess"`
	StartDate          *time.Time      `json:"startDate"`
	EndDate            *time.Time      `json:"endDate"`
	Actors             []Actor         `json:"actors"`
	PreviousActivityID string          `json:"previousActivityId"`
	AcceptActivityID   string          `json:"acceptActivityId"`
	RejectActivityID   string          `json:"rejectActivityId"`
	State              State           `json:"state"`
	ActivityInvokedBy  Actor           `json:"activityInvokedBy"`
	IsAsyncActivity    bool            `json:"isAsyncActivity"`
	PostCondition      PostCondition   `json:"postCondition"`
	PreCondition       PreCondition    `json:"preCondition"`
}

//PreCondition : Activity Preconditions
type PreCondition struct {
}

//PostCondition : Activity Preconditions
type PostCondition struct {
	Targets []TargetIdentifier `json:"targets"`
}

//TargetIdentifier : Target Activity Identifier
type TargetIdentifier struct {
	ConditionExpression string `json:"conditionExpression"`
	TargetActivityID    string `json:"targetActivityId"`
}
