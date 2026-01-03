// auth.go
package handlers

import (
	"encoding/json"
	"net/http"

	// "strconv"
	// "encoding/json"
	// "real-time-forum/internal/database/queries"
	"real-time-forum/internal/database/queries"
	"real-time-forum/internal/models"
	"fmt"
)

// Auth handler here

//receives form data through an HTTP request sent by the frontend (JS)

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

		exists, err := queries.UserOrEmailExists(req.Username, req.Email)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if exists {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Username or email already exists",
		})
		return
	}

	err = queries.CreateUser(req.Username, req.Email, string(req.Password))
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "sessionID",
		Value: req.Username,
		Path: "/",
		HttpOnly: true,
		Secure: false,
		MaxAge: 60*60*24*7,
	})

	fmt.Println("SUCCESSSSS")


		json.NewEncoder(w).Encode(map[string]string{
		"message": "Registration successful",
	})

}

