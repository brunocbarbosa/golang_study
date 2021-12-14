package process_transaction

type transactionDtoInput struct {
	ID        string
	AccountId string
	Amount    float64
}

type transactionDtoOutput struct {
	ID           string
	Status       string
	ErrorMessage string
}
