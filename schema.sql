CREATE TABLE IF NOT EXISTS championship (
  id  INTEGER PRIMARY KEY,
  name text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS team (
  id  INTEGER PRIMARY KEY,
  name text    NOT NULL,
  country text NOT NULL,
  class text NOT NULL,
  championship_id integer,
  FOREIGN KEY (championship_id)
       REFERENCES championship (id) 
);

-- standings will be 'independent' from teams to be able to have denormalized results for all the different type of championships