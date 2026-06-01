-- +goose Up
CREATE TABLE users (
	id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	email TEXT UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE users;
