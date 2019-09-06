package wfmodels

import "time"

//WorkflowProcess : Represent Workflow Process
type WorkflowProcess struct {
	ProcessID                  string     `json:"processId"`
	ProcessName                string     `json:"processName"`
	ProcessCode                string     `json:"processCode"`
	StartDate                  *time.Time `json:"startDate"`
	EndDate                    *time.Time `json:"endDate"`
	SequenceNo                 int        `json:"sequenceNo"`
	NextProcessID              string     `json:"nextProcessId"`
	RootProcessID              string     `json:"rootProcessId"`
	RootActivityID             string     `json:"rootActivityId"`
	IsAsyncProcess             bool       `json:"isAsyncProcessId"`
	ProcessInvokedByActivityID string     `json:"processInvokedByActivityId"`
}
