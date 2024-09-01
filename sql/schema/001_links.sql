-- +goose Up
CREATE TABLE links (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  full_link  VARCHAR(128) NOT NULL,
  short_link VARCHAR(16) UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE links;
