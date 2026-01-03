//auth.go
package models


import (
	"time"
)

//auth structs here


type RegisterRequest struct{
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}


type Cookie struct {
	Username string 
	LoginAt time.Time 
	ExpireAt time.Time 
}

