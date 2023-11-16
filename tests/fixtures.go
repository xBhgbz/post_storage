//go:build integration

package tests

import (
	"storage/pkg/posting"
)

type PostBuilder struct {
	post *posting.Post
}

type CommentBuilder struct {
	comment *posting.Comment
}

func Post() *PostBuilder {
	return &PostBuilder{post: &posting.Post{}}
}

func (p *PostBuilder) Likes(likes int64) *PostBuilder {
	p.post.Likes = likes
	return p
}

func (p *PostBuilder) Text(text string) *PostBuilder {
	p.post.Text = text
	return p
}

func (p *PostBuilder) Pointer() *posting.Post {
	return p.post
}

func (p *PostBuilder) Value() posting.Post {
	return *p.post
}

func (p *PostBuilder) Valid() *PostBuilder {
	return Post().Likes(100).Text("Hello world")
}

func Comment() *CommentBuilder {
	return &CommentBuilder{comment: &posting.Comment{}}
}

func (c *CommentBuilder) Likes(likes int64) *CommentBuilder {
	c.comment.Likes = likes
	return c
}

func (c *CommentBuilder) Text(text string) *CommentBuilder {
	c.comment.Text = text
	return c
}

func (c *CommentBuilder) Value() posting.Comment {
	return *c.comment
}

func (c *CommentBuilder) Pointer() *posting.Comment {
	return c.comment
}

func (c *CommentBuilder) Valid() *CommentBuilder {
	return Comment().Text("simple comment").Likes(228)
}
