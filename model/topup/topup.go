package topup

type TopUpRequest struct {
	Amount int `json:"amount" binding:"required"`
}

type TopUpResponseDB struct {
	TopUpID       string `json:"top_up_id"`
	AmountTopUp   int    `json:"amount_top_up"`
	BalanceBefore int    `json:"balance_before"`
	BalanceAfter  int    `json:"balance_after"`
	CreatedDate   string `json:"created_date"`
}

type TopUpResponse struct {
	Status string      `json:"status"`
	Result TopUpResult `json:"result"`
}

type TopUpResult struct {
	TopUpID       string `json:"top_up_id"`
	AmountTopUp   int    `json:"amount_top_up"`
	BalanceBefore int    `json:"balance_before"`
	BalanceAfter  int    `json:"balance_after"`
	CreatedDate   string `json:"created_date"`
}
