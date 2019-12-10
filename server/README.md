# Redemption Golang API

Clean API to generates redeem links and redeem the balances from Trust Wallet.

## Setup

### Quick start

Deploy it in less than 30 seconds!

### Prerequisite
* [GO](https://golang.org/doc/install) `1.13+`
* Locally running [Redis](https://redis.io/topics/quickstart) or url to remote instance (required for Observer only)

### From Source 

```shell
go get -u github.com/trustwallet/redemption/server
cd $GOPATH/src/github.com/trustwallet/redemption/server

// Make commands
- install     Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
- start       Start API in development mode.
- stop        Stop development mode.
- start-api   Start API in development mode.
- compile     Compile the binary.
- exec        Run given command. e.g; make exec run="go test ./..."
- clean       Clean build files. Runs `go clean` internally.
- test        Run all unit tests.
- fmt         Run `go fmt` for all go files.
```

### Environment Variables

All environment variables for developing are set inside the .env file.

### Docker

Build and run from local Dockerfile:

```shell
docker build -t redeem-heroku .
docker run -p 8399:8399 -p 3000:8080 -d redeem-heroku
```

### Tools

-   Setup MongoDb

```shell
brew install mongodb
```

-   Running in the IDE ( GoLand )

1.  Run
2.  Edit configuration
3.  New Go build configuration
4.  Select `directory` as configuration type
5.  Set `api` as program argument and `-i` as Go tools argument 

### Docs

Swagger API docs provided at path `/swagger/index.html`

#### Updating Docs

- After creating a new route, add comments to your API source code, [See Declarative Comments Format](https://swaggo.github.io/swaggo.io/declarative_comments_format/).
- Download Swag for Go by using:

    `$ go get -u github.com/swaggo/swag/cmd/swag`

- Run the Swag in your Go project root folder.

    `$ swag init`

### Unit Tests

To run the unit tests: `make test`

### Metrics

The application can collect and expose by `expvar's`, metrics about the application healthy and clients and server requests.
Prometheus or another service can collect metrics provided from the `/metrics` endpoint.

To protect the route, you can set the environment variables `METRICS_API_TOKEN`, and this route starts to require the auth bearer token. 
