-- +goose Up
CREATE TABLE links (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  full_link  TEXT NOT NULL,
  short_link TEXT NOT NULL
);

-- +goose Down
DROP TABLE links;
