package handlers

import (
	"fmt"
	"net/http"

	internals "gihub.com/dyxgou/server/internals/db"
	"gihub.com/dyxgou/server/internals/models"
	"gihub.com/dyxgou/server/internals/utils"
)

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
