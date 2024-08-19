package handlers

import (
	"net/http"

	"gihub.com/dyxgou/server/internals/utils"
)

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}
