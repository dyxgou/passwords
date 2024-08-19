package models

type Password struct {
	Id            int    `json:"id"`
	Content       string `json:"content"`
	Sha           string `json:"sha"`
	Rot           string `json:"rot"`
	SecurityLevel string `json:"security_level"`
	UserId        int    `json:"user_id"`
}
