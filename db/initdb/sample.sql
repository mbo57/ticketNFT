USE sample;

DROP TABLE IF EXISTS staff;

CREATE TABLE staff
(
    id       INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    email    VARCHAR(40)  UNIQUE,
    password VARCHAR(300) ,
    name     VARCHAR(40)
);

INSERT INTO staff (email, password, name) VALUES ("test@example.com","password", "佐藤太郎");
INSERT INTO staff (email, password, name) VALUES ("suzuki@example.com", "password", "鈴木一郎");
INSERT INTO staff (email, password, name) VALUES ("user1@example.com", "password", "user1");
INSERT INTO staff (email, password, name) VALUES ("abc@abc", "123", "abc");
