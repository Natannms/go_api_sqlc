CREATE TABLE users (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL
);

CREATE TABLE metadata (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  user_id INTEGER NOT NULL,
  user_type TEXT,
  photo TEXT,
  description TEXT,
  status TEXT,
  address TEXT,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE instrument (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  user_id INTEGER NOT NULL,
  is_active BOOLEAN,
  instrument_name TEXT,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE invites (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  user_id INTEGER NOT NULL,
  remittee TEXT,
  status TEXT,
  invite_type TEXT,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE solicitation (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  user_id INTEGER NOT NULL,
  remittee TEXT,
  status TEXT,
  invite_type TEXT,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE repertoire (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  user_id INTEGER NOT NULL,
  name TEXT,
  is_active BOOLEAN,
  music_list_id INTEGER,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
