// comments.go
package queries

import (
	"fmt"
	"real-time-forum/internal/database"
	"real-time-forum/internal/models"
)

// Comment queries here

func GetPostComments(postID int) ([]models.Comment, error) {
	rows, err := database.DB.Query(`
	SELECT c.id, c.content, u.username, c.created_at
	FROM comments c
	JOIN users u ON c.user_id = u.id
	WHERE c.post_id = ?
	ORDER BY c.created_at DESC
	`, postID)

	if err != nil {
		fmt.Println("error occured getting all comments for this post: ", err)
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment

	for rows.Next() {
		var comment models.Comment

		if err := rows.Scan(&comment.ID, &comment.Content, &comment.Author, &comment.CreatedAt); err != nil {
			continue
		}
		comment.PostID = postID
		comments = append(comments, comment)
	}

	return comments, nil

}
