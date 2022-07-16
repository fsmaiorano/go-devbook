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

func (postRepository posts) FindByID(id uint64) (models.Post, error) {
	var post models.Post

	line, err := postRepository.db.Query("SELECT p.id, p.title, p.content, p.author_id, p.likes, p.created_at, p.updated_at, u.nickname FROM posts p INNER JOIN users u on u.id = p.author_id WHERE p.id = @post_id ", sql.Named("post_id", id))
	if err != nil {
		return post, err
	}

	if line.Next() {
		if err := line.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.UpdatedAt, &post.AuthorNickname); err != nil {
			return post, err
		}
	}

	defer line.Close()

	return post, nil
}

func (postRepository posts) FindAll(userID uint64) ([]models.Post, error) {
	var post models.Post

	lines, err := postRepository.db.Query("SELECT DISTINCT p.id, p.title, p.content, p.author_id, p.likes, p.created_at, p.updated_at, u.nickname FROM posts p INNER JOIN users u on u.id = p.author_id WHERE p.author_id = @user_id order by p.created_at desc", sql.Named("user_id", userID))
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var posts []models.Post
	for lines.Next() {
		if err := lines.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.UpdatedAt, &post.AuthorNickname); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (postRepository posts) Update(id uint64, post models.Post) error {
	statement, err := postRepository.db.Prepare("UPDATE posts SET title = @title, content = @content WHERE id = @id")
	if err != nil {
		return err
	}

	_, err = statement.Exec(sql.Named("id", id), sql.Named("title", post.Title), sql.Named("content", post.Content))
	if err != nil {
		return err
	}

	return nil
}

func (postRepository posts) Delete(id uint64) error {
	statement, err := postRepository.db.Prepare("DELETE FROM posts WHERE id = @id")
	if err != nil {
		return err
	}

	_, err = statement.Exec(sql.Named("id", id))
	if err != nil {
		return err
	}

	return nil
}

func (postRepository posts) FindByUserID(userID uint64) ([]models.Post, error) {
	var post models.Post

	lines, err := postRepository.db.Query("SELECT p.id, p.title, p.content, p.author_id, p.likes, p.created_at, p.updated_at, u.nickname FROM posts p INNER JOIN users u on u.id = p.author_id WHERE p.author_id = @user_id order by p.created_at desc", sql.Named("user_id", userID))
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var posts []models.Post
	for lines.Next() {
		if err := lines.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.UpdatedAt, &post.AuthorNickname); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (postRepository posts) Like(postID uint64) error {
	statement, err := postRepository.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = @id")
	if err != nil {
		return err
	}

	_, err = statement.Exec(sql.Named("id", postID))
	if err != nil {
		return err
	}

	return nil
}

func (postRepository posts) Unlike(postID uint64) error {
	statement, err := postRepository.db.Prepare("UPDATE posts SET likes = likes - 1 WHERE id = @id")
	if err != nil {
		return err
	}

	_, err = statement.Exec(sql.Named("id", postID))
	if err != nil {
		return err
	}

	return nil
}
