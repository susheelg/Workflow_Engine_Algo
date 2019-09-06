package main1

import (
	"WorkflowEngine/datamdl"
	"WorkflowEngine/model"
	"WorkflowEngine/modules/workflow"
	"fmt"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/configmdl"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/loggermdl"
	"go.uber.org/zap/zapcore"
)

func main() {

	loggermdl.Init("logs/Workflow.log", 0, 1, 0, zapcore.DebugLevel)

	//Read config.toml and set values to global config object
	_, configReadError := configmdl.InitConfig(model.GetConfigFilePath(), &model.Config)
	if configReadError != nil {
		loggermdl.LogError("configReadError: ", configReadError)
		return
	}

	mongoInitError := datamdl.InitMongo()
	if mongoInitError != nil {
		loggermdl.LogError("mongoInitError : ", mongoInitError)
		return
	}

	fmt.Println("Welcome to Workflow Engine !!")
	fmt.Println("Mongo in use : ", model.Config.MongoURL+":"+model.Config.MongoPort)
	fmt.Println("DBName : ", model.Config.MongoDBName)

	// mongoClient := datamdl.GetMongoConnection()

	TestFunc()

}

//TestFunc - For testing fucntionalities
func TestFunc() {

	// newScheme := model.Scheme{}
	// newScheme.SchemeName = "Test scheme"
	// newScheme.SequenceNo = 1

	// insertedScheme, insertError := workflow.AddNewScheme(newScheme)
	// if insertError != nil {
	// 	loggermdl.LogError("Insert error : ", insertError)
	// }

	// fmt.Println("Inserted document : ", insertedScheme)
	// schemeObj, errorMsg := workflow.GetSchemeBySchemeID("5d651974ebcc0719f8e8cca6")
	// fmt.Println("Main error : ", errorMsg)
	// fmt.Println("Inside Main : ", schemeObj)

	newProcess := model.Process{}
	newProcess.ProcessCode = "PRC0001"

	newProcess, addProcessError := workflow.AddProcessToScheme("5d651974ebcc0719f8e8cca5", newProcess)
	if addProcessError != nil {
		loggermdl.LogError("Add process : ", addProcessError)
	}

	fmt.Println("Process added : ", newProcess)

}
