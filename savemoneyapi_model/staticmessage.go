package model

const errorWhileFetchingResults = "There was an error while fetching results."
const errorWhileConnectingSQLServer = "There was an error while connecting Server."
const onlyGetTypeIsAllowed = "Only GET type of request is allowed."
const onlyPOSTTypeIsAllowed = "Only POST type of request is allowed."
const invalidHeaderRequest = "invalid header request. we accept only application/json type of requests."
const invalidBodyRequest = "invalid Body request. one of the fields is not correct."
const alreadyExists = "Data that you are trying to insert already exists!"
const invalidQuerystringParameter = "Invalid query string parameter."
const onlyDELETETypeIsAllowed = "Only DELETE type of request is allowed."
const onlyPUTTypeIsAllowed = "Only PUT type of request is allowed."
const userIdDoesNotExists = "UserId does not exists!"
const sqlConnectTimeOut = "Connection to database timed out"

type errorMessage struct {
	Message string
}

type ErrorMessage struct {
	Message error
}

func ReturnSqlTimeOut() errorMessage {
	e := new(errorMessage)
	e.Message = sqlConnectTimeOut
	return *e
}
func ReturnErrorForSQLFetch() errorMessage {
	e := new(errorMessage)
	e.Message = errorWhileFetchingResults
	return *e
}
func ReturnErrorForSQLConnect() errorMessage {
	e := new(errorMessage)
	e.Message = errorWhileConnectingSQLServer
	return *e
}

func OnlyPUTTypeIsAllowed() errorMessage {
	e := new(errorMessage)
	e.Message = onlyPUTTypeIsAllowed
	return *e
}

func OnlyDELETETypeIsAllowed() errorMessage {
	e := new(errorMessage)
	e.Message = onlyDELETETypeIsAllowed
	return *e
}

func OnlyGetTypeIsAllowed() errorMessage {
	e := new(errorMessage)
	e.Message = onlyGetTypeIsAllowed
	return *e
}

func OnlyPOSTTypeIsAllowed() errorMessage {
	e := new(errorMessage)
	e.Message = onlyPOSTTypeIsAllowed
	return *e
}

func InvalidHeaderRequest() errorMessage {
	e := new(errorMessage)
	e.Message = invalidHeaderRequest
	return *e
}

func DataExists() errorMessage {
	e := new(errorMessage)
	e.Message = alreadyExists
	return *e
}

func InvalidQuerystringParameter() errorMessage {
	e := new(errorMessage)
	e.Message = invalidQuerystringParameter
	return *e
}
func InvalidBodyRequest() errorMessage {
	e := new(errorMessage)
	e.Message = invalidBodyRequest
	return *e
}

func UserIdDoesNotExists() errorMessage {
	e := new(errorMessage)
	e.Message = userIdDoesNotExists
	return *e
}

func CustomErrorMessage(errorMessageString string) errorMessage {
	e := new(errorMessage)
	e.Message = errorMessageString
	return *e
}
