DROP TABLE IF EXISTS purchase_items;
DROP TABLE IF EXISTS purchases;
DROP TABLE IF EXISTS price_change;
DROP TABLE IF EXISTS deliveries;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS manufacturers;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS stores;
DROP TABLE IF EXISTS customers;
  
  
CREATE TABLE users (
    id serial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null,
    userName varchar not null,
    seccondName varchar not null,
    isadmin boolean DEFAULT false
  );


CREATE TABLE categories 
(
    category_id SERIAL PRIMARY KEY, 
    category_name VARCHAR(100) NOT NULL
);




CREATE TABLE product_list (
product_id SERIAL PRIMARY KEY, 
FOREIGN KEY(product_id) REFERENCES product (product_id)


)

CREATE TABLE product (
product_id serial  primary key,
product_name varchar not null, 
product_img_path varchar ,
product_price int not null ,
product_discount int DEFAULT 0,
product_description varchar ,
product_compostion varchar not null , 
remaining_products int not null DEFAULT 0,
category_id BIGINT,
FOREIGN KEY (category_id) REFERENCES categories (category_id),
);


CREATE TABLE purchases
(
    purchase_id SERIAL PRIMARY KEY,
    customer_id BIGINT UNSIGNED NOT NULL,
    purchase_date DATETIME NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES users (id),
);

CREATE TABLE purchase_items
(
    purchase_id BIGINT UNSIGNED NOT NULL,
    product_id BIGINT UNSIGNED NOT NULL,
    product_count BIGINT UNSIGNED NOT NULL,
    product_price NUMERIC(9,2) NOT NULL,
    CONSTRAINT PK_PURCHASE_ITEMS PRIMARY KEY (purchase_id, product_id),  
    FOREIGN KEY (product_id) REFERENCES products (product_id),
    FOREIGN KEY (purchase_id) REFERENCES purchases (purchase_id)
);