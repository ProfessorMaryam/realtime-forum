// comments.go
package handlers

import (
	"net/http"
	"strconv"
	"encoding/json"
    "real-time-forum/internal/database/queries"
	"fmt"
)

// Comments handler here


func CommentsHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query().Get("post")
	if q == ""{
		http.Error(w, "missing post id", http.StatusBadRequest)
        return
	}

	fmt.Println("this is query string: ", q)


	   postID, err := strconv.Atoi(q)
    if err != nil {
        http.Error(w, "invalid post id", http.StatusBadRequest)
        return
    }

	comments, err:= queries.GetPostComments(postID)
	if err != nil {
		http.Error(w, "Failed to load comments", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comments)

}