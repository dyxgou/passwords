package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	internals "gihub.com/dyxgou/server/internals/db"
	"gihub.com/dyxgou/server/internals/utils"
	"github.com/go-chi/chi/v5"
)

func GetNameById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	db := internals.InitDB()

	getUserByNameQuery := `SELECT name FROM users WHERE id = ?`

	rows := db.QueryRow(getUserByNameQuery, userId)

	var name string

	err := rows.Scan(&name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithErr(w, 404, fmt.Sprintf("User not found %v", err))
			return
		}
		utils.RespondWithErr(w, 400, fmt.Sprintf("GetUserById Scanning rows error: %v", err))
		return
	}

	type getIdByNameRes struct {
		Username string `json:"user_name"`
	}

	utils.RespondWithJSON(w, 200, getIdByNameRes{
		Username: name,
	})
}
