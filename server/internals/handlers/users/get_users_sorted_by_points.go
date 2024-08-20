package handlers

import (
	"fmt"
	"net/http"

	internals "gihub.com/dyxgou/server/internals/db"
	"gihub.com/dyxgou/server/internals/models"
	"gihub.com/dyxgou/server/internals/utils"
)

func GetUsersSortedByPoints(w http.ResponseWriter, r *http.Request) {
	db := internals.InitDB()

	getUsersSortedByPoints := `SELECT * FROM users ORDER BY points DESC`

	rows, err := db.Query(getUsersSortedByPoints)
	if err != nil {
		utils.RespondWithErr(w, 500, fmt.Sprintf("GetAllPasswords: %v", err))
		return
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User

		err = rows.Scan(
			&u.Id,
			&u.Name,
			&u.Passwords,
			&u.Points,
		)
		if err != nil {
			utils.RespondWithErr(w, 500, fmt.Sprintf("GetAllPasswords: %v", err))
			return
		}

		users = append(users, u)
	}

	utils.RespondWithJSON(w, 200, users)
}
