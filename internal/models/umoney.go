package models

import "time"

type UMoney struct {
	OperationId string `json:"operation_id"`
	NotificationType string `json:"notification_type"`
	Datetime time.Time `json:"datetime"`
	Sha1Hash string `json:"sha1_hash"`
	Sender string `json:"sender"`
	Codepro string `json:"codepro"`
	Currency string `json:"currency"`
	Amount string `json:"amount"`
	WithdrawAmount string `json:"withdraw_amount"`
	Label string `json:"label"`
}
