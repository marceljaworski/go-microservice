# Orders API Webserver with go

API microservice for handling HTTP CRUD requests.

## Features

- chi router and middleware
- redis database
- Graceful shutdown. signint (signal interrupt)

## Build

`p` project name
`docker compose -p postgres up -d`

## Postgres usefull commands

exec container
`docker exec -it postgres bash`

change to postgres user
`su - postgres`

show schema orders
`\d orders`
