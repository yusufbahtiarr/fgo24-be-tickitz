CREATE Table times (
  id SERIAL PRIMARY KEY,
  time TIME NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);