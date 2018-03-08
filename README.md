# goRestExample
GO rest application build on Golang and using Gorilla mux for exposing rest api

Use it like this!

`go run gorest.go`

## Setting up Mysql DB docker container

`docker run --name gorest-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin MYSQL_DATABASE=gorest -d mysql`

### login using mysql clinet

`docker run -it --link gorest-mysql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'`

### create a database

```
create database gorest
use gorest
```

### create table user

`CREATE TABLE users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50) NOT NULL, age INT NOT NULL);`

## Build and Run goRestExample rest application

Install Dependencies

* Get golang. We use 1.9.2.
* Get dep.
* Get yarn.

```
dep ensure
yarn run dev
```

The rest api would be available at http://localhost:1340

## Reference

https://medium.com/@kelvin_sp/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6


