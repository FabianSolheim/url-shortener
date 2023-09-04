-- +goose Up
CREATE TABLE link (
		id SERIAL PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		link TEXT NOT NULL
);

-- +goose Down
DROP TABLE link;