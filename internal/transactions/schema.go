package transactions

type Transaction struct {
	Amount      int    `json:"amount"`
	Operation   string `json:"operation"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
