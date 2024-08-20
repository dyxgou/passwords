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

func GetIdByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	db := internals.InitDB()

	getUserByNameQuery := `SELECT id FROM users WHERE name = ?`

	rows := db.QueryRow(getUserByNameQuery, name)

	var id int64

	err := rows.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithErr(w, 404, fmt.Sprintf("User not found %v", err))
			return
		}
		utils.RespondWithErr(w, 400, fmt.Sprintf("GetUserById Scanning rows error: %v", err))
		return
	}

	type getIdByNameRes struct {
		UserId int64 `json:"user_id"`
	}

	utils.RespondWithJSON(w, 200, getIdByNameRes{
		UserId: id,
	})
}
