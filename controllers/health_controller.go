package controllers

import (
	"encoding/json"
	"net/http"
)

// Ping godoc
// @Summary      Health check
// @Description  Returns OK
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /ping [get]
func Ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
