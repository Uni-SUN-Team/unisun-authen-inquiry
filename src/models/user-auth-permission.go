package models

type UserAuthPermission struct {
	UserId       int     `json:"user_id"`
	TokenVersion int     `json:"tk_version"`
	Iat          float64 `json:"iat"`
	Ext          float64 `json:"ext"`
}
