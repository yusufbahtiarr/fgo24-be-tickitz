CREATE TABLE profile (
  id SERIAL PRIMARY KEY,
  fullname VARCHAR(100),
  phone VARCHAR(15),
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);