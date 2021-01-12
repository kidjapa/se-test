## Test Dev

```shell
go version: 1.15.4
```

## Make
- Run ```make help```
- Compile: ```make compile```

Environment Variables (in .env file)
---------------------

Name | Description | Value | Default |
-----|-----------|---------|--------------|
PORT | Define server port | ```9999``` | ```9090```
GO_ENV | Define the Environment | ```local```, ```homol```, ```prod``` | ```local```

### Make swagger files
1. install `swag`:
```shell
go get github.com/swaggo/swag/cmd/swag
```
2. In root project folder, execute: 
```shell
swag init -g ./main.go
```
3. Access in ```http://localhost:[PORT]/swagger/index.html```

- More info about documentation headers: [aqui](https://github.com/swaggo/swag)

### Run Docker

```shell
docker-compose up --build
```

### Run Tests

```shell
go test ./...
```
