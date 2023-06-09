CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  guid UUID DEFAULT uuid_generate_v4(),
  username VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE UNIQUE INDEX username_uindex ON users(username);

CREATE TABLE IF NOT EXISTS accounts (
  id SERIAL PRIMARY KEY,
  guid UUID DEFAULT uuid_generate_v4(),
  user_id INT REFERENCES users(id),
  account_number VARCHAR NOT NULL,
  balance DECIMAL DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE UNIQUE INDEX account_number_uindex ON accounts(account_number);

CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  guid UUID DEFAULT uuid_generate_v4(),
  sender VARCHAR NOT NULL,
  recipient VARCHAR NOT NULL,
  amount DECIMAL NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
)