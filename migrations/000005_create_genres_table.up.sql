CREATE TABLE genres (
  id SERIAL PRIMARY KEY,
  genre_name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);