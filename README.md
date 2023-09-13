# post-articles-backend

## How to run
### Manually migrate table
use this mysql query
``` sql
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
