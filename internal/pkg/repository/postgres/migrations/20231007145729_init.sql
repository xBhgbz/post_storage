-- +goose Up
-- +goose StatementBegin
CREATE TABLE post (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY NOT NULL,
    text text NOT NULL DEFAULT '',
    likes BIGINT NOT NULL DEFAULT 0,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE comment (
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    text text NOT NULL DEFAULT '',
    likes BIGINT NOT NULL DEFAULT 0,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    post_id BIGINT REFERENCES post (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table comment;
drop table post;
-- +goose StatementEnd
