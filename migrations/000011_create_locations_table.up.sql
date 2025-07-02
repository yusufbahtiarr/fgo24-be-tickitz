CREATE TABLE locations (
  id SERIAL PRIMARY KEY,
  location VARCHAR(50) NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);