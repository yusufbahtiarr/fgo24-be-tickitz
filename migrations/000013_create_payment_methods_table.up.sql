CREATE TABLE payment_methods (
  id SERIAL PRIMARY KEY,
  payment_method VARCHAR(50) NOT NULL,
  image_url VARCHAR(255) NOT NULL
);