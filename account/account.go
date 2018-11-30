package account

type User struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	AccountNumber string `json:"account_number"`
	Name          string `json:"name"`
	Balance       int    `json:"balance"`
}
