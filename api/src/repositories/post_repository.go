package repositories

import (
	"api/src/models"
	"database/sql"
)

type posts struct {
	db *sql.DB
}

func NewRepositoryOfPosts(db *sql.DB) *posts {
	return &posts{db}
}

func (postRepository posts) Create(post models.Post, userID uint64) (uint64, error) {
	statement, err := postRepository.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (@title, @content, @author_id)")
	if err != nil {
		return 0, err
	}

	_, err = statement.Exec(sql.Named("author_id", post.AuthorID), sql.Named("title", post.Title), sql.Named("content", post.Content))
	if err != nil {
		return 0, err
	}

	line, err := postRepository.db.Query("SELECT id FROM posts WHERE author_id = @author_id AND title = @title AND content = @content order by id desc", sql.Named("author_id", post.AuthorID), sql.Named("title", post.Title), sql.Named("content", post.Content))
	if err != nil {
		return 0, err
	}

	if line.Next() {
		if err := line.Scan(&post.ID); err != nil {
			return 0, err
		}
	}

	defer line.Close()

	return post.ID, nil
}
