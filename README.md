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
# server app listening on http://localhost:3000
# web app listening on http://localhost
make
```

### Start server app only
```sh
# server app listening on http://localhost:3000
make serve
```

### Start server app docs
```sh
# server app docs listening on http://localhost:9000/docs
make serve.docs
```

### Start web app only (requires dependencies install)
```sh
# web app listening on http://localhost
make web.serve
```

### Install web app dependencies
```sh
make web.install
```

## Architectural decisions

### Technology stack

#### Backend
I decided to use Go on the backend because, besides being my favorite language, it's a super-fast language as it compiles to machine code and the output binary is very small. Also, it's a strongly, statically typed language that helps developers to catch bugs when compiling. It has a very good standard library that enables developers to build software without the need for external dependencies (that sometimes can lead to security breaches). The concurrency model (goroutines) is probably the easiest one between programming languages which enables services to process heavy workloads quite fast.

#### Frontend
I decided to use React on the frontend because it has a good performance on the client-side and also because it's the one I have more knowledge with.

### Architectur

## Refrences
- [golang project layout standards](https://github.com/golang-standards/project-layout)
