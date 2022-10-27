CREATE TABLE todos (
  id serial PRIMARY KEY,
  title VARCHAR(250) NOT NULL,
  memo VARCHAR(250),
  is_done boolean DEFAULT false NOT NULL,
  due_date TIMESTAMP NOT NULL
);
