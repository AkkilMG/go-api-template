package models

type Basic struct {
	ID     string `json:"id"`
	Detail string `json:"detail"`
	Visits int    `json:"visits"`
}
