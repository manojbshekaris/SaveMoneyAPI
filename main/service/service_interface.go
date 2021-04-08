package service

import (
	"net/http"
)

type IAPIMethods interface {
	//income
	GETTransactionTypes(w http.ResponseWriter, r *http.Request)
	InsertNewIncome(w http.ResponseWriter, r *http.Request)
	GetIncomeByUserId(w http.ResponseWriter, r *http.Request)
	DeleteIncome(w http.ResponseWriter, r *http.Request)
	GetDeletedIncome(w http.ResponseWriter, r *http.Request)
	ReactivateDeletedIncome(w http.ResponseWriter, r *http.Request)
	UpdateBalance(w http.ResponseWriter, r *http.Request)
	//category
	AddCategory(w http.ResponseWriter, r *http.Request)
	GetCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	GetDeletedCategory(w http.ResponseWriter, r *http.Request)
	ReActivateDeletedCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	//Transaction
	AddTransaction(w http.ResponseWriter, r *http.Request)
	GetTransactions(w http.ResponseWriter, r *http.Request)
	HandleSelfTransfer(w http.ResponseWriter, r *http.Request)
}
