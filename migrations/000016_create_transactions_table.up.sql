CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  phone VARCHAR(100) NOT NULL,
  total_payment INT NOT NULL,
  movie_date DATE NOT NULL,
  id_movie INT REFERENCES movies(id),
  id_cinema INT REFERENCES cinemas(id),
  id_time INT REFERENCES times(id),
  id_location INT REFERENCES locations(id),
  id_payment_method INT REFERENCES payment_methods(id),
  id_user INT REFERENCES users(id),
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);