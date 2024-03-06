package transactions

type Transaction struct {
	Amount      int    `json:"valor"`
	Operation   string `json:"tipo"`
	Description string `json:"descricao"`
	CreatedAt   string `json:"realizada_em"`
}
