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

func GetUserPoints(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	db := internals.InitDB()

	getUserPoints := `SELECT points FROM users WHERE id = ?`

	rows := db.QueryRow(getUserPoints, userId)

	var points int

	err := rows.Scan(&points)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithErr(w, 404, fmt.Sprintf("User not found %v", err))
			return
		}
		utils.RespondWithErr(w, 400, fmt.Sprintf("GetUserPoints Scanning rows error: %v", err))
		return
	}

	type userPoints struct {
		Points int `json:"points"`
	}

	utils.RespondWithJSON(w, 200, userPoints{
		Points: points,
	})
}
