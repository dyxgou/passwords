package handlers

import (
	"net/http"

	"gihub.com/dyxgou/server/internals/utils"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithErr(w, 500, "Something went wrong")
}
