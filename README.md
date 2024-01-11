# Orders API Webserver with go

API microservice for handling HTTP CRUD requests.

## Features

- chi router and middleware
- redis database
- Graceful shutdown. signint (signal interrupt)

## Build

`p` project name
`docker compose -p postgres up -d`

## Start app

`go run main.go`

## Postgres usefull commands

exec container
`docker exec -it containerName bash`

Change to postgres user
`su - postgres`

Sign into the database
`psql orders`

List of the available tables
`\d`

Show table orders_db
`\d orders_db`

In this project, the user, db and tables have already been created. But there are these commands to do it:

`CREATE USER postgres_user WITH PASSWORD 'password';`
`CREATE DATABASE my_postgres_db OWNER postgres_user;`
