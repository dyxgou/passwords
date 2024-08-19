package handlers

import (
	"fmt"
	"net/http"

	internals "gihub.com/dyxgou/server/internals/db"
	"gihub.com/dyxgou/server/internals/models"
	"gihub.com/dyxgou/server/internals/utils"
)

func GetAllPasswords(w http.ResponseWriter, r *http.Request) {
	db := internals.InitDB()

	getAllPasswordsQuery := `SELECT id, content, rot, sha, user_id FROM passwords`

	rows, err := db.Query(getAllPasswordsQuery)
	if err != nil {
		utils.RespondWithErr(w, 500, fmt.Sprintf("GetAllPasswords: %v", err))
		return
	}

	defer rows.Close()

	var passwords []models.Password

	for rows.Next() {
		var password models.Password

		err = rows.Scan(
			&password.Id,
			&password.Content,
			&password.Rot,
			&password.Sha,
			&password.UserId,
		)
		if err != nil {
			utils.RespondWithErr(w, 500, fmt.Sprintf("GetAllPasswords: %v", err))
			return
		}

		passwords = append(passwords, password)
	}

	utils.RespondWithJSON(w, 200, passwords)
}
