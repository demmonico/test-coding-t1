CREATE TABLE IF NOT EXISTS users (
   id SERIAL PRIMARY KEY,
   age INT,
   first_name TEXT,
   username TEXT,
   email TEXT UNIQUE NOT NULL
);

INSERT INTO users (age, email, first_name, username)
VALUES (36, 'demmonico@gmail.com', 'Dima', 'demmonico'),
       (63, 'ocinommed@gmail.com', 'Amid', 'ocinommed');