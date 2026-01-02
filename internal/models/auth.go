//auth.go
package models




//auth structs here


type RegisterRequest struct{
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}


