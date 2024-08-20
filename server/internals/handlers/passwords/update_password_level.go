package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	internals "gihub.com/dyxgou/server/internals/db"
	"gihub.com/dyxgou/server/internals/utils"
)

const (
	terrible  = "terrible"
	bad       = "bad"
	mid       = "mid"
	good      = "good"
	excellent = "excellent"
)

type updatePasswordRequest struct {
	PasswordId    int    `json:"password_id"`
	SecurityLevel string `json:"security_level"`
}

func isValidLevel(level string) bool {
	OPTIONS := [5]string{terrible, bad, mid, good, excellent}

	for _, sl := range OPTIONS {
		if sl == level {
			return true
		}
	}

	return false
}

func UpdatePasswordLevel(w http.ResponseWriter, r *http.Request) {
	var b updatePasswordRequest

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		utils.RespondWithErr(w, 400, "Invalid request body")
		return
	}

	if !isValidLevel(b.SecurityLevel) {
		utils.RespondWithErr(w, 400, "Invalid secure level option")
		return
	}

	updatePasswordLevel := `UPDATE passwords SET security_level = ? WHERE id = ?`

	db := internals.InitDB()

	_, err = db.Exec(updatePasswordLevel, b.SecurityLevel, b.PasswordId)
	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("UpdatingPassword : %v", err))
		return
	}

	type updatePasswordResponse struct {
		UpdatedPassword bool `json:"updated_password"`
	}

	utils.RespondWithJSON(w, 200, updatePasswordResponse{
		UpdatedPassword: true,
	})
}
