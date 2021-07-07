# audio mastery service

## Requirements
### Dockerized
- [Docker](https://docs.docker.com/get-docker/) - 20.10.x
- [Docker-compose](https://docs.docker.com/compose/install/) - 1.24.x
### Local
- [Go](https://golang.org/doc/install) - 1.16.x
- [Node](https://nodejs.org/en/download/) - v14.15.x

## Basic Usage
### Start services (requires docker)
```sh
# server app listenning on http://localhost:3000
# web app listenning on http://localhost
make
```
### Start server app only
```sh
# server app listenning on http://localhost:3000
make serve
```
### Start server app docs
```sh
# server app docs listenning on http://localhost:9000/docs
make serve.docs
```
### Start web app only (requires dependencies install)
```sh
# web app listenning on http://localhost
make web.serve
```
### Install web app dependencies
```sh
make web.install
```

## Refrences
- [golang project layout standards](https://github.com/golang-standards/project-layout)
