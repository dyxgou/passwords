package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	internals "gihub.com/dyxgou/server/internals/db"
	"gihub.com/dyxgou/server/internals/models"
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

	_, err = db.Exec(createUserQuery, u.Name)
	if err != nil {
		utils.RespondWithErr(w, 500, fmt.Sprintf("Failed to create the user %s", err))
		return
	}

	type requestResponse struct {
		UserCreated bool `json:"user_created"`
	}

	utils.RespondWithJSON(w, 200, requestResponse{
		UserCreated: true,
	})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := internals.InitDB()
	getAllUsersQuery := "SELECT * FROM users"

	rows, err := db.Query(getAllUsersQuery)
	if err != nil {
		utils.RespondWithErr(w, 500, fmt.Sprintf("GetAllUsers : %v", err))
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.Id, &user.Name, &user.Passwords, &user.Points)
		if err != nil {
			utils.RespondWithErr(w, 500, fmt.Sprintf("Can't scan over the users : %v", err))
			return
		}

		users = append(users, user)
	}

	utils.RespondWithJSON(w, 200, users)
}
