package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	internals "gihub.com/dyxgou/server/internals/db"
	"gihub.com/dyxgou/server/internals/utils"
)

type createPasswordRequest struct {
	Password string `json:"password"`
	UserId   int    `json:"user_id"`
}

type passwords struct {
	rot string
	sha string
}

func (ps *passwords) addRot(rot string) {
	ps.rot = rot
}

func (ps *passwords) addSha(sha string) {
	ps.sha = sha
}

func CreatePassword(w http.ResponseWriter, r *http.Request) {
	var p createPasswordRequest

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		utils.RespondWithErr(w, 400, "Invalid request body")
		return
	}

	db := internals.InitDB()
	pswdch := make(chan utils.HashResults, 2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go utils.HashSHA256(ctx, p.Password, pswdch)
	go utils.HashRot(ctx, p.Password, pswdch)

	var ps passwords

	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			utils.RespondWithErr(w, 408, "The hashing failed or timeout")
			continue
		case result := <-pswdch:
			if result.Err != nil {
				utils.RespondWithErr(
					w,
					408,
					fmt.Sprintf("Failed Hashing with %s : %v", result.HashType, result.Err),
				)
				return
			}

			fmt.Printf("result : %v\n", result)

			switch result.HashType {
			case utils.ROT:
				ps.addRot(result.Hash)
			case utils.SHA256:
				ps.addSha(result.Hash)
			}
		}
	}

	// fmt.Printf("ps : %v\n", ps)

	createPasswordQuery := `INSERT INTO passwords (content, sha, rot, user_id) VALUES (?, ?, ?, ?)`

	_, err = db.Exec(createPasswordQuery, p.Password, ps.sha, ps.rot, p.UserId)
	if err != nil {
		utils.RespondWithErr(w, 500, fmt.Sprintf("Error creating the password : %v", err))
		return
	}

	type createPasswordRes struct {
		CreatedPassword bool `json:"created_password"`
	}

	utils.RespondWithJSON(w, 200, createPasswordRes{
		CreatedPassword: true,
	})
}
