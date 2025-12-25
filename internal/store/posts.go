package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserId    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatesAt string    `json:"update_at"`
	Comments  []Comment `json:"comments"`
	Version   int       `json:"version"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (content,title, user_id, tags)
	VALUES ($1, $2, $3, $4)
	RETURNING id, create_at, update_at
	`
	ctxt, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctxt,
		query,
		post.Content,
		post.Title,
		post.UserId,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatesAt,
	)

	if err != nil {
		fmt.Println(err, "err in create post")
		return err

	}

	return nil

}

func (s *PostStore) GetById(ctx context.Context, id int64) (*Post, error) {
	query := `
		SELECT id, user_id, title, content, create_at, update_at, tags, version FROM posts 
		WHERE id = $1
	`

	var post Post

	ctxt, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctxt,
		query,
		id,
	).Scan(
		&post.ID,
		&post.UserId,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatesAt,
		pq.Array(&post.Tags),
		&post.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil

}

func (s *PostStore) UpdatePost(ctx context.Context, post *Post) error {
	query := `
	UPDATE posts
	SET title = $1, content = $2, version = version + 1
	WHERE id = $3 AND version = $4
	RETURN version
	`

	ctxt, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	if err := s.db.QueryRowContext(ctxt, query, post.Title, post.Content, post.ID, post.Version).Scan(
		&post.Version,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *PostStore) DeleteById(ctx context.Context, id int) error {
	query := `
	DELETE FROM posts 
	WHERE id = $1
	`

	ctxt, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctxt, query, id)

	if err != nil {
		return err
	}

	row, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if row == 0 {
		return ErrNotFound
	}

	return nil

}
