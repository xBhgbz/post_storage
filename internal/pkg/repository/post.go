package repository

import "time"

type Post struct {
	ID        int64     `db:"id"`
	Text      string    `db:"text"`
	Likes     int64     `db:"likes"`
	CreatedAt time.Time `db:"created_at"`
}
