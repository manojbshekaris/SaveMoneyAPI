package model

type StructGetTransactionTypes struct {
	structTransTypeId
	structTransactionTypeName
}
type structTrnAmount struct {
	Amount float64
}
type structTransactionDescription struct {
	Description string
}
type structTransactionTypeName struct {
	TransactionValue string
}
type structTransTypeId struct {
	TransactionTypeId int
}
type transactionDate struct {
	TransactionDate string
}
type structName struct {
	Name string
}
type structCatId struct {
	CategoryId int
}
type structBalance struct {
	Balance float64
}
type structUserId struct {
	IUserId int64
}
type structIncomeId struct {
	IncomeId int64
}
type structCatName struct {
	CategoryName string
}
type StructIncome struct {
	structName
	structBalance
	structUserId
}
type StructSqlIncome struct {
	structName
	structBalance
	structIncomeId
}
type StructUpdateIncome struct {
	structBalance
	structUserId
	structIncomeId
}
type StructSqlManipulateStatus struct {
	Success      bool
	RowsAffected int
}

type StructAddCat struct {
	structUserId
	structCatName
}
type StructUpdtCat struct {
	StructAddCat
	structCatId
}
type StructCatDetails struct {
	structCatId
	structCatName
}

type StructNewTransaction struct {
	structUserId
	structTransTypeId
	structIncomeId
	structCatId
	structTransactionDescription
	Amount float64
	transactionDate
}

type StructGetTransactions struct {
	structUserId
	FromDate string
	ToDate   string
}

type StructTransactionData struct {
	structCatName
	structTransactionTypeName
	structName
	transactionDate
	structTransactionDescription
	structTrnAmount
}

type StructHandlerSelfTransfer struct {
	structUserId
	FromIncomeId int
	ToIncomeId   int
	structTrnAmount
	transactionDate
}
