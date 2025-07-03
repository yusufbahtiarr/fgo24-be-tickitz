CREATE TABLE movies (
  id SERIAL PRIMARY KEY,
  poster_url VARCHAR(255) NOT NULL,
  backdrop_url VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  release_date DATE NOT NULL,
  runtime INT NOT NULL,
  overview TEXT NOT NULL,
  rating DECIMAL(3,1) NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);