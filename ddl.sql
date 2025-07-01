-- STRUCTURE

CREATE TYPE role_type AS ENUM ('Admin', 'User');
CREATE TYPE status_payment_type AS ENUM('Paid', 'Not Paid');
CREATE TYPE status_ticket_type AS ENUM ('In Active', 'Used');

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(100) NOT NULL,
  first_name VARCHAR(100),
  last_name VARCHAR(100),
  phone VARCHAR(15),
  role role_type DEFAULT 'User'
  created_at TIMESTAMP(0) DEFAULT now()
  updated_at TIMESTAMP(0) DEFAULT now()
);

CREATE TABLE sessions (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  token TEXT NOT NULL UNIQUE,
  is_active BOOLEAN DEFAULT TRUE,
  id_user INT REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE casts (
  id SERIAL PRIMARY KEY,
  cast_name VARCHAR(100) NOT NULL
  created_at TIMESTAMP(0) DEFAULT now()
  updated_at TIMESTAMP(0) DEFAULT now()
);

CREATE TABLE genres (
  id SERIAL PRIMARY KEY,
  genre_name VARCHAR(100) NOT NULL
  created_at TIMESTAMP(0) DEFAULT now()
  updated_at TIMESTAMP(0) DEFAULT now()
);

CREATE TABLE directors (
  id SERIAL PRIMARY KEY,
  director_name VARCHAR(100) NOT NULL
  created_at TIMESTAMP(0) DEFAULT now()
  updated_at TIMESTAMP(0) DEFAULT now()
)

CREATE TABLE movies (
  id SERIAL PRIMARY KEY,
  poster_url VARCHAR(255) NOT NULL,
  backdrop_url VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  release_date DATE NOT NULL,
  runtime INT NOT NULL,
  overview TEXT NOT NULL,
  rating DECIMAL(3,1) NOT NULL,
  id_director INT REFERENCES directors(id) ON UPDATE CASCADE ON DELETE SET NULL
  created_at TIMESTAMP(0) DEFAULT now()
  updated_at TIMESTAMP(0) DEFAULT now()
);

CREATE TABLE movie_casts (
  id_movie INT REFERENCES movies(id) ON UPDATE CASCADE ON DELETE CASCADE,
  id_cast INT REFERENCES casts(id) ON UPDATE CASCADE ON DELETE CASCADE
);


CREATE TABLE movie_genres (
  id_movie INT REFERENCES movies(id) ON UPDATE CASCADE ON DELETE CASCADE,
  id_genre INT REFERENCES genres(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE cinemas (
  id SERIAL PRIMARY KEY,
  cinema_name VARCHAR(50) NOT NULL,
  image_url VARCHAR(255) NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);

CREATE TABLE locations (
  id SERIAL PRIMARY KEY,
  location VARCHAR(50) NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);

CREATE Table times (
  id SERIAL PRIMARY KEY,
  time TIME NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);

CREATE TABLE payment_methods (
  id SERIAL PRIMARY KEY,
  payment_method VARCHAR(50) NOT NULL,
  image_url VARCHAR(255) NOT NULL
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  phone VARCHAR(100) NOT NULL,
  virtual_account VARCHAR(50) NOT NULL,
  total_payment INT NOT NULL,
  transaction_date TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  movie_date DATE NOT NULL,
  status_payment status_payment_type DEFAULT 'Not Paid',
  status_ticket status_ticket_type DEFAULT 'In Active',
  id_movie INT REFERENCES movies(id),
  id_cinema INT REFERENCES cinemas(id),
  id_time INT REFERENCES times(id),
  id_location INT REFERENCES locations(id),
  id_payment_method INT REFERENCES payment_methods(id),
  id_user INT REFERENCES users(id),
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);

CREATE TABLE transaction_details(
  id_transaction INT NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
  seat VARCHAR(255) NOT NULL,
  created_at TIMESTAMP(0) DEFAULT now(),
  updated_at TIMESTAMP(0) DEFAULT now()
);

--INDEXING
CREATE INDEX idx_movies_title ON movies (title);

