-- +migrate Up
CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  user_id VARCHAR(10) NOT NULL,
  name VARCHAR(50) NOT NULL,
  category VARCHAR(50) NOT NULL,
  shops VARCHAR(50),
  price DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +migrate Down