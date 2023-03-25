DROP TABLE IF EXISTS Products;
DROP TABLE IF EXISTS Shops;
DROP TABLE IF EXISTS Categories;

CREATE TABLE Categories (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (`id`)
);

CREATE TABLE Shops (
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(255) NOT NULL UNIQUE,
  address     VARCHAR(255) NOT NULL UNIQUE,
  password     VARCHAR(255) NOT NULL,
  hash     VARCHAR(255) NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE Products (
    id INT AUTO_INCREMENT NOT NULL,
    shop_id INT,
    name VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255),
    categories VARCHAR(255),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`shop_id`) REFERENCES Shops(`id`)
);

INSERT INTO Categories (name) VALUES ('Pets'),('Clothing'),('Food');