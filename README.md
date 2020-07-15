# Samproxy Resolution Test

Emulates some of the code in samproxy that adds the machine to the list of peers in redis and then the code that tries to find itself in the list of peers from redis.

## How to run?

`go run main.go`

## Additional options

The default run of the tool will use the default listen address from the samproxy toml file. You can pass in an alternate like so:

`go run main.go 0.0.0.0:8081`
