package api

import (
	"encoding/json"
	"goRestExample/dao"
	"goRestExample/utils"
	"net/http"
)

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dao.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := user.CreateUser(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}
