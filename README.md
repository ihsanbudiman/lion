## Pre-requisites
* Go 1.23.3
* Docker
* Soda v6 for migration
    * `go install github.com/gobuffalo/pop/v6/soda@latest`

## How to run
* run docker compose 
    * `$docker compose up -d`
* run migration
    * `$ make migrate-source-up` 
    * `$ make migrate-destination-up`
* run seed
    * `$ make seed`

## Sync Product Data
you can access sync product endpoint at `http://localhost:8080/sync-product` with POST