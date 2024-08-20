package handlers

import (
	"fmt"
	"net/http"

	internals "gihub.com/dyxgou/server/internals/db"
	"gihub.com/dyxgou/server/internals/models"
	"gihub.com/dyxgou/server/internals/utils"
	"github.com/go-chi/chi/v5"
)

func GetUserPasswords(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	getPasswordsQuery := `SELECT * FROM passwords WHERE user_id = ?`

	db := internals.InitDB()

	rows, err := db.Query(getPasswordsQuery, userId)
	if err != nil {
		utils.RespondWithErr(w, 404, fmt.Sprintf("GetUserPasswords: %v", err))
		return
	}

	var pswds []models.Password

	for rows.Next() {
		var p models.Password
		err := rows.Scan(
			&p.Id,
			&p.Content,
			&p.Rot,
			&p.Sha,
			&p.SecurityLevel,
			&p.UserId,
		)
		if err != nil {
			utils.RespondWithErr(w, 500, fmt.Sprintf("GetUserPasswords: %v", err))
			return
		}

		pswds = append(pswds, p)
	}

	utils.RespondWithJSON(w, 200, pswds)
}
