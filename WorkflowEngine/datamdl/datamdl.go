package datamdl

import (
	"WorkflowEngine/model"
	"strconv"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/dalmdl/coremongo"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/loggermdl"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Client

//InitMongo - Intialized connection with mongodb server
func InitMongo() error {

	// hostName := "WorkflowDB"

	mongoHost := coremongo.MongoHost{}

	mongoHost.HostName = model.Config.HostName
	mongoHost.Server = model.Config.MongoURL
	mongoHost.Port, _ = strconv.Atoi(model.Config.MongoPort)
	mongoHost.Username = model.Config.MongoUsername
	mongoHost.Password = model.Config.MongoPass
	mongoHost.Database = model.Config.MongoDBName
	mongoHost.IsDefault = true

	mongoHostList := []coremongo.MongoHost{mongoHost}
	// mongoHostList = append(mongoHostList, mongoHost)

	initError := coremongo.InitUsingJSON(mongoHostList)
	// initError := coremongo.InitNewSession(mongoHost)

	if initError != nil {
		loggermdl.LogError("datamdl init : ", initError)
		return initError
	}

	newMongoClient, clientError := coremongo.GetMongoConnection(model.Config.HostName)
	if clientError != nil {
		loggermdl.LogError("InitMongo clientError : ", InitMongo)
		return clientError
	}

	mongoClient = newMongoClient

	//Initialize all DAO for DB operations
	InitDAO()

	return nil
}

//GetMongoConnection - Gets mongo client instance
func GetMongoConnection() *mongo.Client {
	return mongoClient
}
