package savemoneyapi_data

import (
	model "moneysaverapi/savemoneyapi_model"
)

type ISqlMethod interface {
	SQLGETTransactionTypes() ([]model.StructGetTransactionTypes, error)
	//Income
	SQLInsertNewIncome(Income model.StructIncome) (model.StructSqlManipulateStatus, error)
	SQLGetIncomeByUserId(UserId int) ([]model.StructSqlIncome, error)
	SQLDeleteIncome(iUserId int, iIncomeId int) (model.StructSqlManipulateStatus, error)
	SQLDELETEDIncomeByUserId(UserId int) ([]model.StructSqlIncome, error)
	SQLReactivateDeletedIncome(iUserId int, iIncomeId int) (model.StructSqlManipulateStatus, error)
	SQLUpdateBalance(model.StructUpdateIncome) (model.StructSqlManipulateStatus, error)
	//Category
	SQLAddCategory(Cat model.StructAddCat) (model.StructSqlManipulateStatus, error)
	SQLGetCategory(UserId int) ([]model.StructCatDetails, error)
	SQLDeleteCategory(UserId int, Category int) (model.StructSqlManipulateStatus, error)
	SQLGetDeletedCategory(iUserId int) ([]model.StructCatDetails, error)
	SQLReActivateDeletedCategory(UserId int, TransactionCategoryId int) (model.StructSqlManipulateStatus, error)
	SQLUpdateCategory(model.StructUpdtCat) (model.StructSqlManipulateStatus, error)
	//Transaction
	SQLAddTransaction(trn model.StructNewTransaction) (model.StructSqlManipulateStatus, error)
	SQLGetTransactions(trnReq model.StructGetTransactions) ([]model.StructTransactionData, error)
	SQLHandleSelfTransfer(trn model.StructHandlerSelfTransfer) (model.StructSqlManipulateStatus, error)
}
