package model

type Customer struct {
	AccountId int64  `json:"account_id,omitempty"`
	Name      string `json:"name,omitempty"`
}
