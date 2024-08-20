package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	internals "gihub.com/dyxgou/server/internals/db"
	"gihub.com/dyxgou/server/internals/models"
	"gihub.com/dyxgou/server/internals/utils"
	"github.com/go-chi/chi/v5"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	db := internals.InitDB()

	getUserByIdQuery := `SELECT * FROM users WHERE id = ?`

	rows := db.QueryRow(getUserByIdQuery, userId)

	var u models.User

	err := rows.Scan(&u.Id, &u.Name, &u.Passwords, &u.Points)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithErr(w, 404, fmt.Sprintf("User not found %v", err))
			return
		}
		utils.RespondWithErr(w, 500, fmt.Sprintf("GetUserById Scanning rows error: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, u)
}
