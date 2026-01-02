// auth.go
package handlers

import (
	"encoding/json"
	"net/http"
	// "strconv"
	// "encoding/json"
	// "real-time-forum/internal/database/queries"
	"real-time-forum/internal/models"
	"fmt"
)

// Auth handler here

//still have not registered in main.go

//receives form data through an HTTP request sent by the front end JS to the Golang backend

func RegisterHandler(w http.ResponseWriter, r *http.Request){

		w.Header().Set("Content-Type", "application/json")

			if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	fmt.Println(req.Username)
	fmt.Println(req.Password)
	fmt.Println(req.Email)




}