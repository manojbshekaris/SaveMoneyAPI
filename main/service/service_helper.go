package service

import (
	"encoding/json"
	utility "moneysaverapi"
	model "moneysaverapi/savemoneyapi_model"
	"net/http"
	"strings"
)

func JsonWebResponse(objectToConvertAsJson interface{}, statusCode int, w http.ResponseWriter) bool {
	w.Header().Set("Content-Type", "application/json")
	results, err := json.Marshal(objectToConvertAsJson)
	if err != nil {
		utility.WriteToLog(err, "****catastrophic failure")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Message": "catastrophic failure has occurred!"}`))
		return false
	} else {
		w.WriteHeader(statusCode)
		w.Write(results)
		return false
	}
}

func HandleSQLRequestErrors(err error, w http.ResponseWriter) bool {
	if err != nil {
		utility.WriteToLog(err)
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return JsonWebResponse(model.DataExists(), http.StatusConflict, w)
		}
		if strings.Contains(strings.ToLower(err.Error()), "1452") {
			return JsonWebResponse(model.UserIdDoesNotExists(), http.StatusConflict, w)
		}
		if strings.Contains(strings.ToLower(err.Error()), "1045") {
			return JsonWebResponse(model.ReturnErrorForSQLConnect(), http.StatusInternalServerError, w)
		}
		if strings.Contains(strings.ToLower(err.Error()), "timeout") {
			return JsonWebResponse(model.ReturnSqlTimeOut(), http.StatusInternalServerError, w)
		}
		//custom errors created at sql.
		if strings.Contains(strings.ToLower(err.Error()), "1644") {
			return JsonWebResponse(model.CustomErrorMessage(strings.Split(err.Error(), "Error 1644: ")[1]), http.StatusInternalServerError, w)
		}
		return JsonWebResponse(model.ReturnErrorForSQLFetch(), http.StatusInternalServerError, w)
	}
	return true
}

func HandleSQLStatus(stats model.StructSqlManipulateStatus, w http.ResponseWriter) bool {
	if stats.RowsAffected > 0 && stats.Success {
		return JsonWebResponse(stats, http.StatusAccepted, w)
	} else if stats.RowsAffected == 0 && !stats.Success {
		return JsonWebResponse(stats, http.StatusAlreadyReported, w)
	}
	return true
}

func HandleQueryStringErr(err error, w http.ResponseWriter) bool {
	if err != nil {
		utility.WriteToLog(err)
		return JsonWebResponse(model.InvalidQuerystringParameter(), http.StatusBadRequest, w)
	}
	return true
}

func PreprocessRequests(acceptedVerb string, w http.ResponseWriter, r *http.Request) bool {
	acceptedVerb = strings.ToUpper(acceptedVerb)
	switch acceptedVerb {
	case "GET":
		if acceptedVerb != strings.ToUpper(r.Method) {
			return JsonWebResponse(model.OnlyGetTypeIsAllowed(), http.StatusMethodNotAllowed, w)
		}
	case "POST":
		if acceptedVerb != strings.ToUpper(r.Method) {
			return JsonWebResponse(model.OnlyPOSTTypeIsAllowed(), http.StatusMethodNotAllowed, w)
		}
		if acceptedVerb == strings.ToUpper(r.Method) && r.Header.Get("Content-Type") != "application/json" {
			return JsonWebResponse(model.InvalidHeaderRequest(), http.StatusUnsupportedMediaType, w)
		}
	case "DELETE":
		if acceptedVerb != strings.ToUpper(r.Method) {
			return JsonWebResponse(model.OnlyDELETETypeIsAllowed(), http.StatusMethodNotAllowed, w)
		}
	case "PUT":
		if acceptedVerb != strings.ToUpper(r.Method) {
			return JsonWebResponse(model.OnlyPUTTypeIsAllowed(), http.StatusMethodNotAllowed, w)
		}
	default:
		//this logic returns bool iff we have matching http verb from request.
		if acceptedVerb == strings.ToUpper(r.Method) {
			return true
		} else {
			return JsonWebResponse(model.CustomErrorMessage("Invalid HTTP verb."), http.StatusMethodNotAllowed, w)
		}
	}
	return true
}

func ExtractStructFromJSONisOk(infce interface{}, r *http.Request, w http.ResponseWriter) bool {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&infce)
	if err != nil {
		utility.WriteToLog(err)
		return JsonWebResponse(model.CustomErrorMessage(err.Error()), http.StatusBadRequest, w)
	}
	return true
}
