package handlers

import "net/http"

const (
	levelHorrible = "horrible"
	levelBad      = "bad"
	levelMid      = "mid"
	levelGood     = "good"
	levelExcelent = "excellent"
)

type editPasswordRequest struct {
	PasswordId    int    `json:"password_id"`
	SecurityLevel string `json:"security_level"`
}

func EditPasswordLevel(w http.ResponseWriter, r *http.Request) {
}
