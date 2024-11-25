package entity

type Customer struct {
	Id   string    `json:"id"`
	Name string    `json:"name"`
	Email string   `json:"email"`
	Password string `json:"password,omitempty"`
	Balance float64 `json:"balance"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}