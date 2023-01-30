package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

// Currency представляет код валюты
type Currency string

// Category представляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д.)
type Category string

// Status представляет собой статус платежа.
type Status string

// Код валюты
const(
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// Предоределение статус платежей
const(
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
) 

// PAN представляет номер карты
type PAN string

// Payment представляет информацию о платеже.
type Payment struct{
	ID int
	Amount Money
	Category Category
	Status Status
}

// Card представляет информацию о платёжной карте
type Card struct {
	ID			int
	PAN			PAN
	Balance		Money // использовали Money
	MinBalance	Money // использовали Money
	Currency	Currency
	Color		string
	Name		string
	Active		bool
}


// uiifyuf