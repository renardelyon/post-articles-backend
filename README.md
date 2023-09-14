# post-articles-backend

## How to run

### Manually migrate table

use this mysql query

```sql
CREATE DATABASE IF NOT EXISTS article;

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
```

### Database migration

- For database migration first create database named `article` using mysql terminal
- initialize repo using

```Makefile
make init
```

- tidy up dependencies using

```Makefile
make tidy
```

- migrate table using

```Makefile
make migration-up
```

### Running the service

- create `config.ini` file using `config.example.ini` as template
- If there are no `go.mod` run this command

```
make init && make tidy
```

- to run the service, run this command

```
make run
```

## API Documentation

API documentation can be seen in this postman [link](https://www.postman.com/descent-module-technologist-74020729/workspace/post-articles-backend/request/19990992-a31920c2-6d8d-470a-af3f-20ffa14e16f0)
