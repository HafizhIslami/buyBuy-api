CREATE TABLE carts (
  id INT NOT NULL PRIMARY KEY,
  order int NOT NULL,
  user_id int REFERENCES users (id) ON DELETE CASCADE,
  product_id int REFERENCES products (id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN  KEY (product_id) REFERENCES products(id)
);