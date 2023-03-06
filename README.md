# ACME PROJECT

This project is an implementation of an authentication api

## Requirements
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Golang](https://go.dev/doc/install) 
- [Migrate](https://github.com/golang-migrate/migrate)
- [Make](https://community.chocolatey.org/packages/make) only windows users this guide
- [Swagg go](https://github.com/swaggo/swag) (not required)


## Description

The purpose of the project is to create a user interface where they register and then signIn and signOut

## How to use it

**Installing and provisioning the db**

In the first instance it is recommended to use docker to raise the service and the necessary dependencies

The project requires a postgresql database to insert the sales records to be processed, so follow the recommendations and install make and run the command
```
make postgres
```

Intall migrate for created new database and run the migrations with the comand

for created new database this comand
```
make createdb
```

for run migrations
```
make migrateup
```

**NOTE**
if you already have installed a local postgresql database higher than 12.x ignore these steps,
and use the sql found in migrations to create the table in a database named store

## Export environment variables.
```
export $(cat .env.example | grep -v ^# | xargs)
```
**NOTE**
only run local, for the docker using -e NAME_ENVIROMENT=$VALUE

## Run with golang instaled

stand at the root of the project
and run

```
go run ./cmd/server/main.go
```

## Using Docker
to build image
```
docker build -t acme-api .
```

to run container
```
docker run --name acme acme-api
```

to run container if the db postgresql is containirized
```
docker run --publish 8080:8080 -e DB_HOST=host.docker.internal  --name acme acme-api
```

## Swagger docs

If the app configuration is in dev, you can review the service documentation by accessing

http://localhost:8080/api/acme/docs/index.html

