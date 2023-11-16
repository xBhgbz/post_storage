package repository

import "time"

type Comment struct {
	ID        int64     `db:"id"`
	PostID    int64     `db:"post_id"`
	Text      string    `db:"text"`
	Likes     int64     `db:"likes"`
	CreatedAt time.Time `db:"created_at"`
}
