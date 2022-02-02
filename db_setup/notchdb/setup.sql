CREATE DATABASE notch;
GO

USE notch
CREATE TABLE usersproduct (
    id INT  PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    price INT NOT NULL,
);

CREATE TABLE cart (
    id INT  PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    totalprice INT ,
);

CREATE TABLE productCart (
    productId INT NOT NULL,
    cartId INT NOT NULL,
    FOREIGN KEY (productId) REFERENCES usersproduct(id),
    FOREIGN KEY (cartId) REFERENCES cart(id)
);

INSERT INTO usersproduct (
    id,
    name, 
    description, 
    price
    ) 
    VALUES 
    (
        1, 
        'Product 1', 
        'Description 1', 
        10.0
    ),
    (
        2, 
        'Product 2', 
        'Description 2', 
        20.0
    ),
    (
        3, 
        'Product 3', 
        'Description 3', 
        30.0
    ),
    (
        4, 
        'Product 4', 
        'Description 4', 
        40.0
    ),
    (
        5, 
        'Product 5', 
        'Description 5', 
        50.0
    ),
    (
        6, 
        'Product 6', 
        'Description 6', 
        60.0
    ),
    (
        7, 
        'Product 7', 
        'Description 7', 
        70.0
    ),
    (
        8, 
        'Product 8', 
        'Description 8', 
        80.0
    ),
    (
        9, 
        'Product 9', 
        'Description 9', 
        90.0
    ),
    (
        10, 
        'Product 10', 
        'Description 10', 
        100.0
    );
    
GO
