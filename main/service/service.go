package service

import (
	utility "moneysaverapi"
	sql "moneysaverapi/savemoneyapi_data"
	model "moneysaverapi/savemoneyapi_model"
	"net/http"
	"strconv"
)

type StructApi struct{}

var ApiService IAPIMethods = new(StructApi)
var sqlMethods sql.ISqlMethod = sql.Methods

func (api StructApi) GETTransactionTypes(w http.ResponseWriter, r *http.Request) {
	if PreprocessRequests("GET", w, r) {
		trArray, err := sqlMethods.SQLGETTransactionTypes()
		if HandleSQLRequestErrors(err, w) {
			if trArray != nil && len(trArray) > 1 {
				JsonWebResponse(trArray, http.StatusOK, w)
			} else {
				JsonWebResponse(trArray, http.StatusNoContent, w)
			}
		}
	}
}

func (api StructApi) InsertNewIncome(w http.ResponseWriter, r *http.Request) {
	structIncome := new(model.StructIncome)
	if PreprocessRequests("POST", w, r) {
		if ExtractStructFromJSONisOk(&structIncome, r, w) {
			if structIncome.IUserId <= 0 || structIncome.Name == "" {
				utility.WriteToLog("structIncome :", structIncome)
				JsonWebResponse(model.InvalidBodyRequest(), http.StatusBadRequest, w)
				return
			} else {
				stats, err := sqlMethods.SQLInsertNewIncome(*structIncome)
				if HandleSQLRequestErrors(err, w) {
					HandleSQLStatus(stats, w)
				}
			}
		}
	}
}

func (api StructApi) GetIncomeByUserId(w http.ResponseWriter, r *http.Request) {
	if PreprocessRequests("GET", w, r) {
		userId, err := strconv.Atoi(r.URL.Query().Get("UserId"))
		if err != nil {
			JsonWebResponse(model.InvalidQuerystringParameter(), http.StatusBadRequest, w)
			return
		} else {
			IncArray, err := sqlMethods.SQLGetIncomeByUserId(userId)
			if HandleSQLRequestErrors(err, w) {
				if len(IncArray) == 0 {
					JsonWebResponse(nil, http.StatusNoContent, w)
					return
				} else {
					JsonWebResponse(IncArray, http.StatusOK, w)
					return
				}
			}
		}
	}
}

func (api StructApi) DeleteIncome(w http.ResponseWriter, r *http.Request) {
	if PreprocessRequests("DELETE", w, r) {
		userId, err := strconv.Atoi(r.URL.Query().Get("UserId"))
		if HandleQueryStringErr(err, w) {
			incomeId, err := strconv.Atoi(r.URL.Query().Get("IncomeId"))
			if HandleQueryStringErr(err, w) {
				if err != nil {
					JsonWebResponse(model.InvalidQuerystringParameter(), http.StatusBadRequest, w)
					return
				} else {
					stats, err := sqlMethods.SQLDeleteIncome(userId, incomeId)
					if HandleSQLRequestErrors(err, w) {
						HandleSQLStatus(stats, w)
					}
				}
			}
		}
	}
}

func (api StructApi) GetDeletedIncome(w http.ResponseWriter, r *http.Request) {
	if PreprocessRequests("GET", w, r) {
		userId, err := strconv.Atoi(r.URL.Query().Get("UserId"))
		if HandleQueryStringErr(err, w) {
			IncArray, err := sqlMethods.SQLDELETEDIncomeByUserId(userId)
			if HandleSQLRequestErrors(err, w) {
				if err == nil {
					if len(IncArray) == 0 {
						JsonWebResponse(nil, http.StatusNoContent, w)
						return
					} else {
						JsonWebResponse(IncArray, http.StatusOK, w)
						return
					}
				}
			}
		}
	}
}

func (api StructApi) ReactivateDeletedIncome(w http.ResponseWriter, r *http.Request) {
	if PreprocessRequests("PUT", w, r) {
		userId, err := strconv.Atoi(r.URL.Query().Get("UserId"))
		if HandleQueryStringErr(err, w) {
			incomeId, err := strconv.Atoi(r.URL.Query().Get("IncomeId"))
			if HandleQueryStringErr(err, w) {
				stats, err := sqlMethods.SQLReactivateDeletedIncome(userId, incomeId)
				if HandleSQLRequestErrors(err, w) {
					HandleSQLStatus(stats, w)
				}
			}
		}
	}
}

func (api StructApi) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	structUpdtInc := new(model.StructUpdateIncome)
	if PreprocessRequests("POST", w, r) {
		if ExtractStructFromJSONisOk(&structUpdtInc, r, w) {
			stats, err := sqlMethods.SQLUpdateBalance(*structUpdtInc)
			if HandleSQLRequestErrors(err, w) {
				HandleSQLStatus(stats, w)
			}
		}
	}
}

func (api StructApi) AddCategory(w http.ResponseWriter, r *http.Request) {
	cat := new(model.StructAddCat)
	if PreprocessRequests("POST", w, r) {
		if ExtractStructFromJSONisOk(&cat, r, w) {
			stats, err := sqlMethods.SQLAddCategory(*cat)
			if HandleSQLRequestErrors(err, w) {
				HandleSQLStatus(stats, w)
			}
		}
	}
}

func (api StructApi) GetCategory(w http.ResponseWriter, r *http.Request) {
	if PreprocessRequests("GET", w, r) {
		userId, err := strconv.Atoi(r.URL.Query().Get("UserId"))
		if HandleQueryStringErr(err, w) {
			IncArray, err := sqlMethods.SQLGetCategory(userId)
			if HandleSQLRequestErrors(err, w) {
				if err == nil {
					if len(IncArray) == 0 {
						JsonWebResponse(nil, http.StatusNoContent, w)
						return
					} else {
						JsonWebResponse(IncArray, http.StatusOK, w)
						return
					}
				}
			}
		}
	}
}

func (api StructApi) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	if PreprocessRequests("DELETE", w, r) {
		userId, err := strconv.Atoi(r.URL.Query().Get("UserId"))
		if HandleQueryStringErr(err, w) {
			categoryId, err := strconv.Atoi(r.URL.Query().Get("CategoryId"))
			if HandleQueryStringErr(err, w) {
				if err != nil {
					JsonWebResponse(model.InvalidQuerystringParameter(), http.StatusBadRequest, w)
					return
				} else {
					stats, err := sqlMethods.SQLDeleteCategory(userId, categoryId)
					if HandleSQLRequestErrors(err, w) {
						HandleSQLStatus(stats, w)
					}
				}
			}
		}
	}
}

func (api StructApi) GetDeletedCategory(w http.ResponseWriter, r *http.Request) {
	if PreprocessRequests("GET", w, r) {
		userId, err := strconv.Atoi(r.URL.Query().Get("UserId"))
		if HandleQueryStringErr(err, w) {
			IncArray, err := sqlMethods.SQLGetDeletedCategory(userId)
			if HandleSQLRequestErrors(err, w) {
				if err == nil {
					if len(IncArray) == 0 {
						JsonWebResponse(nil, http.StatusNoContent, w)
						return
					} else {
						JsonWebResponse(IncArray, http.StatusOK, w)
						return
					}
				}
			}
		}
	}
}

func (api StructApi) ReActivateDeletedCategory(w http.ResponseWriter, r *http.Request) {
	if PreprocessRequests("PUT", w, r) {
		userId, err := strconv.Atoi(r.URL.Query().Get("UserId"))
		if HandleQueryStringErr(err, w) {
			CategoryId, err := strconv.Atoi(r.URL.Query().Get("CategoryId"))
			if HandleQueryStringErr(err, w) {
				stats, err := sqlMethods.SQLReActivateDeletedCategory(userId, CategoryId)
				if HandleSQLRequestErrors(err, w) {
					HandleSQLStatus(stats, w)
				}
			}
		}
	}
}

func (api StructApi) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	structUpdtcat := new(model.StructUpdtCat)
	if PreprocessRequests("POST", w, r) {
		if ExtractStructFromJSONisOk(&structUpdtcat, r, w) {
			stats, err := sqlMethods.SQLUpdateCategory(*structUpdtcat)
			if HandleSQLRequestErrors(err, w) {
				HandleSQLStatus(stats, w)
			}
		}
	}
}

func (api StructApi) AddTransaction(w http.ResponseWriter, r *http.Request) {
	trn := new(model.StructNewTransaction)
	if PreprocessRequests("POST", w, r) {
		if ExtractStructFromJSONisOk(&trn, r, w) {
			stats, err := sqlMethods.SQLAddTransaction(*trn)
			if HandleSQLRequestErrors(err, w) {
				HandleSQLStatus(stats, w)
			}
		}
	}
}

func (api StructApi) GetTransactions(w http.ResponseWriter, r *http.Request) {
	t := new(model.StructGetTransactions)
	if PreprocessRequests("POST", w, r) {
		if ExtractStructFromJSONisOk(&t, r, w) {
			res, err := sqlMethods.SQLGetTransactions(*t)
			if HandleSQLRequestErrors(err, w) {
				if err == nil {
					if len(res) == 0 {
						JsonWebResponse(nil, http.StatusNoContent, w)
						return
					} else {
						JsonWebResponse(res, http.StatusOK, w)
						return
					}
				}
			}
		}
	}
}

func (api StructApi) HandleSelfTransfer(w http.ResponseWriter, r *http.Request) {
	selfTr := new(model.StructHandlerSelfTransfer)
	if PreprocessRequests("POST", w, r) {
		if ExtractStructFromJSONisOk(&selfTr, r, w) {
			stats, err := sqlMethods.SQLHandleSelfTransfer(*selfTr)
			if HandleSQLRequestErrors(err, w) {
				HandleSQLStatus(stats, w)
			}
		}
	}
}
