package clients

type Client struct {
	ID           int `json:"id"`
	AccountLimit int `json:"limite"`
	Balance      int `json:"total"`
}
