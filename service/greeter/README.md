# Greeter Service

This is the Greeter service

## Configuration

- FQDN: mkit.service.greeter
- Type: service
- Alias: greeter

## Usage

A Makefile is included for convenience

Build the binary

```bash
make build TARGET=greeter TYPE=service
# then run with custom env
MICRO_SERVER_ADDRESS=0.0.0.0:9999 CONFIGOR_ENV_PREFIX=APP APP_FEATURES_TLS_ENABLED=true ./build/greeter-service
```

Run the service

```bash
make run-greeter
# or
go run service/greeter/main.go service/greeter/plugin.go
```

Build a docker image

```bash
make docker TARGET=greeter TYPE=service VERSION=v0.1.1
```

### Test

```bash
# start the server on fixed port
make run-greeter ARGS="--server_address=:8080 --broker_address=:10001"

# test with grpc cli
grpcurl -plaintext -proto service/greeter/proto/greeter/greeter.proto list
grpcurl -plaintext -proto service/greeter/proto/greeter/greeter.proto describe
grpcurl -plaintext -proto service/greeter/proto/greeter/greeter.proto -d '{"name": "sumo"}' localhost:8080  mkit.service.greeter.Greeter/Hello
# testing via micro-cli
micro --client=grpc call --metadata trans-id=1234 mkit.service.greeter Greeter.Hello  '{"name": "John"}'

# start REST gateway
micro api --enable_rpc=true

# testing via rest proxy
curl --request POST \
--url http://localhost:8080/rpc \
--header 'accept: application/json' \
--header 'content-type: application/json' \
--data '{"service": "mkit.service.greeter", "method": "Greeter.Hello","request": {"name": "sumo"}}'
```
