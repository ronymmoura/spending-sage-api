CREATE TABLE monthly_expected_payments (
  id bigserial PRIMARY KEY,
  name varchar NOT NULL,
  amount int NOT NULL CHECK (amount > 0),
  day int2 NOT NULL,
  user_id bigserial NOT NULL
);

CREATE TABLE expected_payments (
  id bigserial PRIMARY KEY,
  name varchar NOT NULL,
  amount int NOT NULL CHECK (amount > 0),
  date date NOT NULL,
  month_id bigserial NOT NULL
);

ALTER TABLE monthly_expected_payments ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE expected_payments ADD FOREIGN KEY (month_id) REFERENCES months (id);