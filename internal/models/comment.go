// comment.go
package models

// Comment model here

import (
	"time"
)


type Comment struct{
	 ID        int       `json:"id"`
	 Content string    `json:"content"`
	 PostID int 		`json:"post_id"`
	 Username string 			`json:"username"`
	 CreatedAt time.Time 		`json:"created_at"`
}