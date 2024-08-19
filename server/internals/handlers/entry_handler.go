package handlers

import (
	"net/http"

	"gihub.com/dyxgou/server/internals/utils"
)

func EntryHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Message string `json:"message"`
	}

	utils.RespondWithJSON(w, 200, response{
		Message: "We're rolling baby!",
	})
}
