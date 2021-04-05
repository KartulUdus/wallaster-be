## Requirements

* [Golang](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/get-docker/) (For postgres database)

## Postgres setup

Run the following command to launch a postgres docker container
```bash
$ docker run --name postgres-db --publish 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_USER=mysecretuser -d postgres:latest
```

## Run the app
To run the app itself run
```bash
go run main.go
```

You should get an output confirming database migration, seeding and `Listening to port 8080`

## Run tests
TBD