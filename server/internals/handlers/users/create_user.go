package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	internals "gihub.com/dyxgou/server/internals/db"
	"gihub.com/dyxgou/server/internals/utils"
)

type createUserRequest struct {
	Name string `json:"name"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u createUserRequest

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		utils.RespondWithErr(w, 400, "Invalid request body")
		return
	}

	if len(u.Name) < 5 {
		utils.RespondWithErr(w, 400, "The name should be greater than 4 chars")
		return
	}

	db := internals.InitDB()

	createUserQuery := `INSERT INTO users (name) VALUES (?)`

	result, err := db.Exec(createUserQuery, u.Name)
	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("User already exists : %s", err))
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		utils.RespondWithErr(w, 500, fmt.Sprintf("LastInsertedId error: %v", err))
		return
	}

	type requestResponse struct {
		UserId int64 `json:"user_id"`
	}

	utils.RespondWithJSON(w, 200, requestResponse{
		UserId: id,
	})
}
