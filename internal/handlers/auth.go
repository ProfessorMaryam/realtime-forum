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
	"real-time-forum/internal/cookie"
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

		exists, err := queries.UserExists(req.Username, req.Email)
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

	cookie , err:= queries.AddSession(req.Email)

	http.SetCookie(w, &cookie)

	fmt.Println("SUCCESSSSS REGISTERING")

		json.NewEncoder(w).Encode(map[string]string{
		"message": "Registration successful",
	})

}

func LoginHandler(w http.ResponseWriter, r *http.Request){
			w.Header().Set("Content-Type", "application/json")
if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil{
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	exists, err := queries.EmailExists(req.Email)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if !exists {
			w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"User does not exist. Please Register",
		})
		return
	}

	cookie , err:= queries.AddSession(req.Email)

	http.SetCookie(w, &cookie)

	fmt.Println("SUCCESSSSS LOGIN")


		json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
	})

}


func LogoutHandler(w http.ResponseWriter, r *http.Request){
			w.Header().Set("Content-Type", "application/json")
if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

		c, err := r.Cookie("sessionID")
		if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	sessionError := queries.DeletePastSessions(c.Value); if sessionError != nil{
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	cookie.DeleteCookie(w,r)


			json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})




}