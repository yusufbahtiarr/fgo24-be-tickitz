CREATE TABLE movie_directors (
  id_movie INT REFERENCES movies(id) ON UPDATE CASCADE ON DELETE CASCADE,
  id_director INT REFERENCES directors(id) ON UPDATE CASCADE ON DELETE CASCADE
);