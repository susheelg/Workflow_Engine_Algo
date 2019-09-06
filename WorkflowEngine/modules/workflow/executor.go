package workflow

//IWorkflow - Interface for executing activities
type IWorkflow interface {
	ExecuteAactivity() (map[string]interface{}, error)
}
