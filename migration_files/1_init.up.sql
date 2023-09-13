CREATE TABLE posts (
    id int NOT NULL AUTO_INCREMENT,
    title varchar(200),
    content text,
    category varchar(100),
    created_date timestamp,
    updated_date timestamp,
    status varchar(100),
    PRIMARY KEY(id)
);
