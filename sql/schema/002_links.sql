-- +goose Up
ALTER TABLE links 
ADD COLUMN counter INTEGER NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE links DROP COLUMN counter;
