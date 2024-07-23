CREATE TABLE users (
  id bigserial PRIMARY KEY,
  email varchar NOT NULL UNIQUE,
  full_name varchar NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (NOW())
);

CREATE TABLE origins (
  id bigserial PRIMARY KEY,
  name varchar NOT NULL,
  type varchar NOT NULL
);

CREATE TABLE categories (
  id bigserial PRIMARY KEY,
  name varchar NOT NULL
);

CREATE TABLE months (
  id bigserial PRIMARY KEY,
  user_id bigserial NOT NULL,
  date date NOT NULL
);

CREATE TABLE month_entries (
  id bigserial PRIMARY KEY,
  month_id bigserial NOT NULL,
  name varchar NOT NULL,
  due_date date NOT NULL,
  pay_date date NOT NULL,
  amount int NOT NULL CHECK (amount > 0),
  owner varchar NOT NULL,
  origin_id bigserial NOT NULL,
  category_id bigserial NOT NULL
);

CREATE TABLE fixed_entries (
  id bigserial PRIMARY KEY,
  user_id bigserial NOT NULL,
  name varchar NOT NULL,
  due_date date NOT NULL,
  pay_day date NOT NULL,
  amount int NOT NULL CHECK (amount > 0),
  owner varchar NOT NULL,
  origin_id bigserial NOT NULL,
  category_id bigserial NOT NULL
);

CREATE TABLE fixed_entry_payment_history (
  id bigserial PRIMARY KEY,
  entry_id bigserial NOT NULL,
  amount int NOT NULL CHECK (amount > 0),
  date timestamptz NOT NULL
);

CREATE INDEX ON users (email);

ALTER TABLE months ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE month_entries ADD FOREIGN KEY (month_id) REFERENCES months (id);

ALTER TABLE month_entries ADD FOREIGN KEY (origin_id) REFERENCES origins (id);

ALTER TABLE month_entries ADD FOREIGN KEY (category_id) REFERENCES categories (id);

COMMENT ON COLUMN "month_entries"."pay_date" IS 'Date that is planned to pay this entry';

ALTER TABLE fixed_entries ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE fixed_entries ADD FOREIGN KEY (origin_id) REFERENCES origins (id);

ALTER TABLE fixed_entries ADD FOREIGN KEY (category_id) REFERENCES categories (id);

ALTER TABLE fixed_entry_payment_history ADD FOREIGN KEY (entry_id) REFERENCES fixed_entries (id);

