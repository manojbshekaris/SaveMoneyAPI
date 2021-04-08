package savemoneyapi_data

import (
	"database/sql"
	utility "moneysaverapi"
	"moneysaverapi/main/config"
	model "moneysaverapi/savemoneyapi_model"

	_ "github.com/go-sql-driver/mysql"
)

type StructSQL struct{}

var Methods ISqlMethod = new(StructSQL)

func (structSql StructSQL) SQLGETTransactionTypes() ([]model.StructGetTransactionTypes, error) {
	var trnTypes []model.StructGetTransactionTypes
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return trnTypes, err
	}
	resRows, err := dataBaseObj.Query("CALL GETTransactionTypes")
	if err != nil {
		utility.WriteToLog(err)
	} else {
		for resRows.Next() {
			trnTypeTemp := new(model.StructGetTransactionTypes)
			resRows.Scan(&trnTypeTemp.TransactionTypeId, &trnTypeTemp.TransactionValue)
			trnTypes = append(trnTypes, *trnTypeTemp)
		}
	}

	defer dataBaseObj.Close()
	return trnTypes, err
}

func (structSql StructSQL) SQLInsertNewIncome(Income model.StructIncome) (model.StructSqlManipulateStatus, error) {
	status := model.StructSqlManipulateStatus{RowsAffected: 0, Success: false}
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	result, err := dataBaseObj.Exec(`CALL InsertNewIncome(?, ?, ?, "NEW INSERT","API"  )`, Income.Name, Income.Balance, Income.IUserId)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		utility.WriteToLog(err)
	} else {
		if nums > 0 {
			status.RowsAffected = int(nums)
			status.Success = true
		}
	}
	return status, err
}

func (structSQL StructSQL) SQLGetIncomeByUserId(UserId int) ([]model.StructSqlIncome, error) {
	var Incomes []model.StructSqlIncome
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return Incomes, err
	}
	resRows, err := dataBaseObj.Query("CALL GetIncomeByUserId(?)", UserId)
	if err != nil {
		utility.WriteToLog(err)
		return Incomes, err
	} else {
		for resRows.Next() {
			IncomeTemp := new(model.StructSqlIncome)
			resRows.Scan(&IncomeTemp.IncomeId, &IncomeTemp.Name, &IncomeTemp.Balance)
			Incomes = append(Incomes, *IncomeTemp)
		}
	}
	defer dataBaseObj.Close()
	return Incomes, err
}

func (structSQL StructSQL) SQLDeleteIncome(iUserId int, iIncomeId int) (model.StructSqlManipulateStatus, error) {
	var status model.StructSqlManipulateStatus
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	result, err := dataBaseObj.Exec(`CALL DeleteIncome(?, ?, "DELETE INCOME", "API" )`, iUserId, iIncomeId)
	if err != nil {
		utility.WriteToLog(err)
		utility.WriteToLog("iUserId : ", iUserId, "iIncomeId", iIncomeId)
		return status, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		utility.WriteToLog(err)
	} else {
		if nums > 0 {
			status.RowsAffected = int(nums)
			status.Success = true
		}
	}
	defer dataBaseObj.Close()
	return status, err
}

func (structSQL StructSQL) SQLDELETEDIncomeByUserId(UserId int) ([]model.StructSqlIncome, error) {
	var Incomes []model.StructSqlIncome
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return Incomes, err
	}
	resRows, err := dataBaseObj.Query("CALL GetDeletedIncome(?)", UserId)
	if err != nil {
		utility.WriteToLog(err)
		return Incomes, err
	} else {
		for resRows.Next() {
			IncomeTemp := new(model.StructSqlIncome)
			resRows.Scan(&IncomeTemp.IncomeId, &IncomeTemp.Name, &IncomeTemp.Balance)
			Incomes = append(Incomes, *IncomeTemp)
		}
	}
	defer dataBaseObj.Close()
	return Incomes, err
}

func (structSQL StructSQL) SQLReactivateDeletedIncome(iUserId int, iIncomeId int) (model.StructSqlManipulateStatus, error) {
	var status model.StructSqlManipulateStatus
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	result, err := dataBaseObj.Exec(`CALL ReactivateDeletedIncome(?, ?, "REACTIVATE INCOME", "API" )`, iUserId, iIncomeId)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		utility.WriteToLog(err)
	} else {
		if nums > 0 {
			status.RowsAffected = int(nums)
			status.Success = true
		}
	}
	defer dataBaseObj.Close()
	return status, err
}

func (structSQL StructSQL) SQLUpdateBalance(strctIncm model.StructUpdateIncome) (model.StructSqlManipulateStatus, error) {
	var status model.StructSqlManipulateStatus
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	result, err := dataBaseObj.Exec(`CALL UpdateBalance(?, ?, ?, "BALANCE UPDATE", "API")`, strctIncm.IUserId, strctIncm.IncomeId, strctIncm.Balance)
	if err != nil {
		utility.WriteToLog(err)
		utility.WriteToLog("iUserId : ", strctIncm.IUserId, "iIncomeId : ", strctIncm.IncomeId, "fBalance : ", strctIncm.Balance)
		return status, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		utility.WriteToLog(err)
	} else {
		if nums > 0 {
			status.RowsAffected = int(nums)
			status.Success = true
		}
	}
	defer dataBaseObj.Close()
	return status, err
}

func (structSQL StructSQL) SQLAddCategory(Cat model.StructAddCat) (model.StructSqlManipulateStatus, error) {
	var status model.StructSqlManipulateStatus
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	result, err := dataBaseObj.Exec(`CALL AddCategory(?, ?, "API", "NEW INSERT")`, Cat.IUserId, Cat.CategoryName)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		utility.WriteToLog(err)
	} else {
		if nums > 0 {
			status.RowsAffected = int(nums)
			status.Success = true
		}
	}
	return status, err
}

func (structSQL StructSQL) SQLGetCategory(UserId int) ([]model.StructCatDetails, error) {
	list := []model.StructCatDetails{}
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return list, err
	}
	resRows, err := dataBaseObj.Query("CALL GetCategory(?)", UserId)
	if err != nil {
		utility.WriteToLog(err)
		return list, err
	} else {
		for resRows.Next() {
			catTemp := new(model.StructCatDetails)
			resRows.Scan(&catTemp.CategoryId, &catTemp.CategoryName)
			list = append(list, *catTemp)
		}
	}
	defer dataBaseObj.Close()
	return list, err
}

func (structSQL StructSQL) SQLDeleteCategory(UserId int, Category int) (model.StructSqlManipulateStatus, error) {
	var status model.StructSqlManipulateStatus
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	result, err := dataBaseObj.Exec(`CALL DeleteCategory(?, ?, "API", "DELETE CATEGORY")`, UserId, Category)
	if err != nil {
		utility.WriteToLog(err)
		utility.WriteToLog("iUserId : ", UserId, "Category", Category)
		return status, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		utility.WriteToLog(err)
	} else {
		if nums > 0 {
			status.RowsAffected = int(nums)
			status.Success = true
		}
	}
	defer dataBaseObj.Close()
	return status, err
}

func (structSQL StructSQL) SQLGetDeletedCategory(iUserId int) ([]model.StructCatDetails, error) {
	{
		list := []model.StructCatDetails{}
		dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
		if err != nil {
			utility.WriteToLog(err)
			return list, err
		}
		resRows, err := dataBaseObj.Query("CALL GetDeletedCategory(?)", iUserId)
		if err != nil {
			utility.WriteToLog(err)
			return list, err
		} else {
			for resRows.Next() {
				catTemp := new(model.StructCatDetails)
				resRows.Scan(&catTemp.CategoryId, &catTemp.CategoryName)
				list = append(list, *catTemp)
			}
		}
		defer dataBaseObj.Close()
		return list, err
	}
}

func (structSQL StructSQL) SQLReActivateDeletedCategory(UserId int, TransactionCategoryId int) (model.StructSqlManipulateStatus, error) {
	var status model.StructSqlManipulateStatus
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	result, err := dataBaseObj.Exec(`CALL ReactivateDeletedCategory(?, ?, "API", "REACTIVATE CATEGORY")`, UserId, TransactionCategoryId)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		utility.WriteToLog(err)
	} else {
		if nums > 0 {
			status.RowsAffected = int(nums)
			status.Success = true
		}
	}
	defer dataBaseObj.Close()
	return status, err
}

func (structSQL StructSQL) SQLUpdateCategory(cat model.StructUpdtCat) (model.StructSqlManipulateStatus, error) {
	var status model.StructSqlManipulateStatus
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	result, err := dataBaseObj.Exec(`CALL UpdateCategory(?, ?, ?, "CATEGORY NAME UPDATE", "API")`, cat.IUserId, cat.CategoryId, cat.CategoryName)
	if err != nil {
		utility.WriteToLog(err)
		utility.WriteToLog("iUserId : ", cat.IUserId, "CategoryId : ", cat.CategoryId, "Category Name : ", cat.CategoryName)
		return status, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		utility.WriteToLog(err)
	} else {
		if nums > 0 {
			status.RowsAffected = int(nums)
			status.Success = true
		}
	}
	defer dataBaseObj.Close()
	return status, err
}

func (structSQL StructSQL) SQLAddTransaction(trn model.StructNewTransaction) (model.StructSqlManipulateStatus, error) {
	var status model.StructSqlManipulateStatus
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	result, err := dataBaseObj.Exec(`CALL InsertNewTransaction(?, ?, ?, ?, ?, ?, ?, "NEW INSERT", "API")`, trn.IUserId, trn.TransactionTypeId, trn.IncomeId, trn.CategoryId, trn.Description, trn.TransactionDate, trn.Amount)
	if err != nil {
		utility.WriteToLog(err)
		utility.WriteToLog("trn object: ", trn)
		return status, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		utility.WriteToLog(err)
	} else {
		if nums > 0 {
			status.RowsAffected = int(nums)
			status.Success = true
		}
	}
	defer dataBaseObj.Close()
	return status, err
}

func (structSQL StructSQL) SQLGetTransactions(trnReq model.StructGetTransactions) ([]model.StructTransactionData, error) {
	list := []model.StructTransactionData{}
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return list, err
	}
	resRows, err := dataBaseObj.Query("CALL GetTransactions(?,?,?)", trnReq.IUserId, trnReq.FromDate, trnReq.ToDate)
	if err != nil {
		utility.WriteToLog(err)
		return list, err
	} else {
		for resRows.Next() {
			t := new(model.StructTransactionData)
			resRows.Scan(&t.CategoryName, &t.TransactionValue, &t.Name, &t.TransactionDate, &t.Description, &t.Amount)
			list = append(list, *t)
		}
	}
	defer dataBaseObj.Close()
	return list, err
}

func (structSQL StructSQL) SQLHandleSelfTransfer(trn model.StructHandlerSelfTransfer) (model.StructSqlManipulateStatus, error) {
	var status model.StructSqlManipulateStatus
	dataBaseObj, err := sql.Open("mysql", config.GetConfigDetails().MySql_Connection_String)
	if err != nil {
		utility.WriteToLog(err)
		return status, err
	}
	result, err := dataBaseObj.Exec(`CALL HandleSelfTransfer(?,?,?,?,?)`, trn.IUserId, trn.FromIncomeId, trn.ToIncomeId, trn.Amount, trn.TransactionDate)
	if err != nil {
		utility.WriteToLog(err)
		utility.WriteToLog("trn object: ", trn)
		return status, err
	}
	nums, err := result.RowsAffected()
	if err != nil {
		utility.WriteToLog(err)
	} else {
		if nums > 0 {
			status.RowsAffected = int(nums)
			status.Success = true
		}
	}
	defer dataBaseObj.Close()
	return status, err
}
