# golang-api-template
A few golang REST endpoints to be used as a starter for new projects

### Start a local redis container
$ docker pull redis
$ docker run --name redis-test-instance -p 6379:6379 -d redis

### Run the app locally
$ go run main.go
