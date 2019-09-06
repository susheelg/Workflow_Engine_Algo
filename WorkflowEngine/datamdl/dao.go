package datamdl

import (
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/dalmdl/coremongo"
)

var (
	//SchemeDAO - For DB operation on Scheme collection
	SchemeDAO *coremongo.MongoDAO

	//ProcessDAO - For DB operation on Process collection
	ProcessDAO *coremongo.MongoDAO

	//ActivityDAO - For DB operation on Activity collection
	ActivityDAO *coremongo.MongoDAO

	//StateDAO - For DB operation on State collection
	StateDAO *coremongo.MongoDAO

	//ActorDAO - For DB operation on Actor collection
	ActorDAO *coremongo.MongoDAO
)

//InitDAO - Associate DAOs with respective collections
//Create new DAO here for new collection
func InitDAO() {
	SchemeDAO = coremongo.GetMongoDAO("Scheme")
	ProcessDAO = coremongo.GetMongoDAO("Process")
	ActivityDAO = coremongo.GetMongoDAO("Activity")
	StateDAO = coremongo.GetMongoDAO("State")
	ActorDAO = coremongo.GetMongoDAO("Actor")
}
