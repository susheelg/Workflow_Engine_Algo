package workflow

import (
	"WorkflowEngine/datamdl"
	"WorkflowEngine/model"
	"errors"
	"strings"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/loggermdl"
)

//AddProcessToScheme - Add a new process to scheme
func AddProcessToScheme(schemeID string, newProcess model.Process) (model.Process, error) {

	schemeID = strings.TrimSpace(schemeID)

	if schemeID == "" {
		loggermdl.LogError("AddProcessToScheme : schemeID required")
		return newProcess, errors.New("AddProcessToScheme : schemeID required")
	}

	//Find if scheme exists
	_, getSchemeError := GetSchemeBySchemeID(schemeID)
	if getSchemeError != nil {
		loggermdl.LogError("AddProcessToScheme getSchemeError : ", getSchemeError)
		return newProcess, getSchemeError
	}

	newProcess.SchemeID = schemeID

	_, saveError := datamdl.ProcessDAO.SaveData(newProcess)
	if saveError != nil {
		loggermdl.LogError("AddProcessToScheme saveError : ", saveError)
		return newProcess, saveError
	}

	return newProcess, nil

}
