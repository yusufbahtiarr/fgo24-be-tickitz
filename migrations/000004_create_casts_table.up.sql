CREATE TABLE casts (
  id SERIAL PRIMARY KEY,
  cast_name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);
