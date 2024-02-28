package clients

type Client struct {
	ID           int `json:"id"`
	AccountLimit int `json:"account_limit"`
	Balance      int `json:"balance"`
}
