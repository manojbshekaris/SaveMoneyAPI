package service

import (
	"moneysaverapi/main/config"
	"net/http"
)

func StartAPIServer(am IAPIMethods) {
	http.HandleFunc("/gettransactiontypes", am.GETTransactionTypes)
	http.HandleFunc("/Income/InsertNewIncome", am.InsertNewIncome)
	http.HandleFunc("/Income/GetIncomeByUserId", am.GetIncomeByUserId)
	http.HandleFunc("/Income/DeleteIncome", am.DeleteIncome)
	http.HandleFunc("/Income/GetDeletedIncome", am.GetDeletedIncome)
	http.HandleFunc("/Income/ReactivateDeletedIncome", am.ReactivateDeletedIncome)
	http.HandleFunc("/Income/UpdateBalance", am.UpdateBalance)
	http.HandleFunc("/Category/AddCategory", am.AddCategory)
	http.HandleFunc("/Category/GetCategory", am.GetCategory)
	http.HandleFunc("/Category/DeleteCategory", am.DeleteCategory)
	http.HandleFunc("/Category/GetDeletedCategory", am.GetDeletedCategory)
	http.HandleFunc("/Category/ReActivateDeletedCategory", am.ReActivateDeletedCategory)
	http.HandleFunc("/Category/UpdateCategory", am.UpdateCategory)
	http.HandleFunc("/Transaction/AddTransaction", am.AddTransaction)
	http.HandleFunc("/Transaction/GetTransactions", am.GetTransactions)
	http.HandleFunc("/Transaction/HandleSelfTransfer", am.HandleSelfTransfer)
	http.ListenAndServe(":"+config.GetConfigDetails().Port, nil)
}
