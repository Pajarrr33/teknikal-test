package request

type TransactionRequest struct {
	CustomerId string `json:"customer_id"`
	MerchantId string `json:"merchant_id"`
	Amount     float64    `json:"amount"`
}