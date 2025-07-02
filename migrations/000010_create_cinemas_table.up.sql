CREATE TABLE cinemas (
  id SERIAL PRIMARY KEY,
  cinema_name VARCHAR(50) NOT NULL,
  image_url VARCHAR(255) NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);