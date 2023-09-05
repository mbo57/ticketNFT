USE sample;

DROP TABLE IF EXISTS staff;
DROP TABLE IF EXISTS cast;
DROP TABLE IF EXISTS eventcategory;
DROP TABLE IF EXISTS event;

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



CREATE TABLE cast
(
    id   INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100)
);

INSERT INTO cast (name) VALUES ("木村花子");
INSERT INTO cast (name) VALUES ("佐藤太郎");
INSERT INTO cast (name) VALUES ("鈴木次郎");


CREATE TABLE eventcategory
(
    id   INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100)
);

INSERT INTO eventcategory (name) VALUES ("スポーツ");
INSERT INTO eventcategory (name) VALUES ("音楽ライブ");


CREATE TABLE event
(
    id              INT           NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name            VARCHAR(100),
    img             VARCHAR(100),
    date            DATE,
    venue           VARCHAR(100),
    castid          INT           REFERENCES cast (id),
    eventcategoryid INT           REFERENCES eventcategory (id),
    description     VARCHAR(500)
);

INSERT INTO event
(
    name,
    img,
    date,
    venue,
    castid,
    eventcategoryid,
    description
)
VALUES
(
    "佐藤誕生祭",
    "/img/test.png",
    "2023-03-10",
    "東京ドーム",
    2,
    2,
    "佐藤誕生ライブ"
), 
(
    "鈴木誕生祭",
    "/img/test.png",
    "2023-09-12",
    "東京ドーム",
    3,
    2,
    "鈴木誕生ライブ"
);


CREATE TABLE users
(
    id              CHAR(50)            NOT NULL PRIMARY KEY ,
    name            VARCHAR(100),
    email           VARCHAR(100)        NOT NULL,
    password        VARCHAR(100)        NOT NULL

);
