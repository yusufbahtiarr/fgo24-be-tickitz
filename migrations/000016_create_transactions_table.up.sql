CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  phone VARCHAR(100) NOT NULL,
  virtual_account VARCHAR(50) NOT NULL,
  total_payment INT NOT NULL,
  transaction_date TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  movie_date DATE NOT NULL,
  status_payment status_payment_type DEFAULT 'not paid',
  status_ticket status_ticket_type DEFAULT 'in active',
  id_movie INT REFERENCES movies(id),
  id_cinema INT REFERENCES cinemas(id),
  id_time INT REFERENCES times(id),
  id_location INT REFERENCES locations(id),
  id_payment_method INT REFERENCES payment_methods(id),
  id_user INT REFERENCES users(id),
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);