package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	internals "gihub.com/dyxgou/server/internals/db"
	"gihub.com/dyxgou/server/internals/utils"
)

func getPoints(securityLevel string) (int, error) {
	points := make(map[string]int)

	points["pending"] = 0
	points["terrible"] = -2
	points["bad"] = -1
	points["mid"] = 0
	points["good"] = 1
	points["excellent"] = 2

	if point, exists := points[securityLevel]; exists {
		return point, nil
	}

	return 0, errors.New("Invalid security level")
}

type addUserPointsRequest struct {
	PasswordId    int    `json:"password_id"`
	SecurityLevel string `json:"security_level"`
}

func AddUserPoints(w http.ResponseWriter, r *http.Request) {
	const PASSWORD_PENDING = "pending"
	var pr addUserPointsRequest

	fmt.Println(pr)
	err := json.NewDecoder(r.Body).Decode(&pr)
	if err != nil {
		utils.RespondWithJSON(w, 400, fmt.Sprintf("Parsing JSON: %v", err))
		return
	}

	points, err := getPoints(pr.SecurityLevel)
	if err != nil {
		utils.RespondWithErr(w, 400, "Invalid securityLevel")
		return
	}

	db := internals.InitDB()
	getPasswordQuery := `SELECT security_level, user_id FROM passwords WHERE id = ?`
	row := db.QueryRow(getPasswordQuery, pr.PasswordId)

	type password struct {
		SecurityLevel string
		UserId        string
	}

	var p password

	err = row.Scan(&p.SecurityLevel, &p.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithErr(w, 404, fmt.Sprintf("Password not found %v", err))
			return
		}
		utils.RespondWithErr(w, 400, fmt.Sprintf("GetUserById Scanning rows error: %v", err))
		return
	}
	type addedUserPoints struct {
		PointsAdded bool `json:"points_added"`
	}

	if p.SecurityLevel != PASSWORD_PENDING {
		utils.RespondWithJSON(w, 200, addedUserPoints{
			PointsAdded: true,
		})
		return
	}

	// updateUserPasswords := `UPDATE users SET points = points + 1 WHERE id = ?`
	updateUserPoints := `UPDATE users SET points = points + ? WHERE id = ?`

	stmt, err := db.Prepare(updateUserPoints)
	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("AddUserPoints Prepare : %v", err))
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(points, p.UserId)
	if err != nil {
		utils.RespondWithErr(w, 400, fmt.Sprintf("AddUserPoints Exec: %v", err))
		return
	}

	utils.RespondWithJSON(w, 200, addedUserPoints{
		PointsAdded: true,
	})
}
