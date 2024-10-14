package transfer

import "time"

type TransferRequest struct {
	TargetUser string `json:"target_user"`
	Amount     int    `json:"amount"`
	Remarks    string `json:"remarks"`
}

type Transfer struct {
	TransferID    string    `json:"transfer_id"`
	Amount        int       `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore int       `json:"balance_before"`
	BalanceAfter  int       `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date"`
}

type TransferResponse struct {
	Status string          `json:"status"`
	Result TransferDetails `json:"result"`
}

type TransferDetails struct {
	TransferID    string    `json:"transfer_id"`
	Amount        int       `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore int       `json:"balance_before"`
	BalanceAfter  int       `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date"`
}
