CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  age INT NOT NULL
);

Create table Orders (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  quantity INT NOT NULL,
  price DECIMAL(10,2) NOT NULL
  FOREIGN KEY (user_id) REFERENCES users(id)
);

