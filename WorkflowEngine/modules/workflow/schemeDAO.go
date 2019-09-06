package workflow

import (
	"WorkflowEngine/datamdl"
	"WorkflowEngine/model"
	"encoding/json"
	"errors"
	"strings"

	"github.com/tidwall/gjson"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/loggermdl"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddNewScheme - Add new scheme to DB
func AddNewScheme(newScheme model.Scheme) (model.Scheme, error) {

	insertedID, insertError := datamdl.SchemeDAO.SaveData(newScheme)
	if insertError != nil {
		loggermdl.LogError("AddNewScheme insertError : ", insertError)
		return newScheme, insertError
	}

	newScheme.SchemeID = insertedID

	return newScheme, nil
}

//GetSchemeBySchemeID - Get a scheme from DB when schemeID is provided
func GetSchemeBySchemeID(schemeID string) (model.Scheme, error) {

	scheme := model.Scheme{}

	schemeID = strings.TrimSpace(schemeID)
	if schemeID == "" {
		loggermdl.LogError("GetSchemeBySchemeID : schemeID required")
		return scheme, errors.New("GetSchemeBySchemeID : schemeID required")
	}

	//Adding parameters for query
	queryParameters := make(map[string]interface{})

	objID, hexError := primitive.ObjectIDFromHex(schemeID)
	if hexError != nil {
		loggermdl.LogError("GetSchemeBySchemeID hexError : ", hexError)
		return scheme, hexError
	}
	queryParameters["_id"] = objID

	queryResult, queryError := datamdl.SchemeDAO.GetData(queryParameters)

	if queryError != nil {
		loggermdl.LogError("GetSchemeBySchemeID queryError : ", queryError)
		return scheme, queryError
	}

	if !gjson.Get(queryResult.Get("0").String(), "schemeName").Exists() {
		return scheme, errors.New("GetSchemeBySchemeID : Scheme '" + schemeID + "' not found")
	}

	unmarshalError := json.Unmarshal([]byte(queryResult.Get("0").String()), &scheme)
	if unmarshalError != nil {
		loggermdl.LogError("GetSchemeBySchemeID unmarshalError: ", unmarshalError)
		return scheme, unmarshalError
	}

	return scheme, nil

}
