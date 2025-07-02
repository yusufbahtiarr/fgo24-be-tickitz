CREATE TABLE directors (
  id SERIAL PRIMARY KEY,
  director_name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);