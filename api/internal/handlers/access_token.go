package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/stephen-pp/yata/api/db"
	"github.com/stephen-pp/yata/api/internal/models"
)

type AccessTokenRequest struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func GenerateAccessToken(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Secret-Token") != os.Getenv("SECRET_TOKEN") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req AccessTokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	accessToken := models.NewAccessToken(db.GetDB(), req.Username, req.ID)

	err = accessToken.Save()
	if err != nil {
		http.Error(w, "Failed to save access token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AccessTokenResponse{AccessToken: accessToken.Token})
}
