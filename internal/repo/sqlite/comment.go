package sqlite

import (
	"fmt"
	"forum/models"
)

func (s *Sqlite) CommentPost(form models.CommentForm) error {
	op := "sqlite.CommentPost"
	stmt := `INSERT INTO Comments (post_id, user_id, content, created) VALUES(?, ?, ?, CURRENT_TIMESTAMP)`
	_, err := s.db.Exec(stmt, form.PostID, form.UserID, form.Content)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Sqlite) GetCommentsByPostID(postID string) (*[]models.Comment, error) {
	const query = `SELECT c.id, c.post_id, c.user_id, c.created, c.content, c.like, c.dislike, u.name 
	FROM comments c 
	JOIN users u ON c.user_id = u.id 
	WHERE c.id = ?`
	rows, err := s.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Created, &comment.Content, &comment.Like, &comment.Dislike, &comment.UserName)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return &comments, nil
}

// like system

func (s *Sqlite) CheckReactionComment(form models.ReactionForm) (bool, bool, error) {

	// Check if the user has already liked/disliked the post
	var isExists bool
	checkQuery := `SELECT EXISTS(SELECT is_like FROM Comment_User_Like WHERE user_id = ? AND comment_id = ?)`
	err := s.db.QueryRow(checkQuery, form.UserID, form.ID).Scan(&isExists)
	if err != nil {
		return false, false, err
	}
	var dbLike bool
	if isExists {
		checkQuery = `SELECT is_like FROM Comment_User_Like WHERE user_id = ? AND comment_id = ?`
		err = s.db.QueryRow(checkQuery, form.UserID, form.ID).Scan(&dbLike)
		if err != nil {
			return false, false, err
		}
	}

	return isExists, dbLike, nil
}

func (s *Sqlite) AddReactionComment(form models.ReactionForm) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Insert like/dislike
	insertQuery := `INSERT INTO Comment_User_Like (user_id, comment_id, is_like) VALUES (?, ?, ?)`
	_, err = tx.Exec(insertQuery, form.UserID, form.ID, form.Reaction)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Update Post like/dislike count
	updateQuery := ""
	if form.Reaction {
		updateQuery = `UPDATE Comments SET like = like + 1 WHERE id = ?`
	} else {
		updateQuery = `UPDATE Comments SET dislike = dislike + 1 WHERE id = ?`
	}
	_, err = tx.Exec(updateQuery, form.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *Sqlite) DeleteReactionComment(form models.ReactionForm, isLike bool) error {
	tx, err := s.db.Begin()
	if err != nil {
		fmt.Println("here")
		return err
	}

	// delete the like/dislike
	deleteQuery := `DELETE FROM Comment_User_Like WHERE user_id = ? AND comment_id = ?`
	_, err = tx.Exec(deleteQuery, form.UserID, form.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// decrement the like or dislike
	updateQuery := ""
	if isLike {
		updateQuery = `UPDATE Comments SET like = like - 1 WHERE id = ? AND like > 0`
	} else {
		updateQuery = `UPDATE Comments SET dislike = dislike - 1  WHERE id = ? AND dislike > 0`
	}
	_, err = tx.Exec(updateQuery, form.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
