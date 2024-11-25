package entity

type Transaction struct {
	Id string `json:"id"`
	CustomerId string `json:"customer_id"`
	MerchantId string `json:"merchant_id"`
	Amount float64 `json:"amount"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}