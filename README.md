# taco 🌮🌮🌮
The next generation repository system for DLSS
![taco](https://user-images.githubusercontent.com/92044/34897877-016a4e36-f7b6-11e7-80e3-4edecfb2f89d.gif)

## Swagger API

This configuration is for AWS API Gateway.  It was retrieved by going to the API, selecting the "prod" under "Stages" and then doing "Export" and selecting "Export as Swagger + API Gateway Extensions"

## Go Local Development Setup

1. Install go (grab binary from here or use `brew install go` on Mac OSX).
2. Setup your Go workspace (where your Go code, binaries, etc. are kept together. TO DO: add info on Go workspace FYI.):
    ```bash
    $ mkdir -p ~/go
    $ export GOPATH=~/go
    $ export PATH=~/go/bin:$PATH
    $ cd ~/go
    ```
    Your Go code repositories will reside within `~/go/src/...` in the `$GOPATH`. Name these paths to avoid library clash, for example TACO Go code could be in `~/go/src/github.com/sul-dlss-labs/taco`. This should be where your Github repository resides too.
3. Change into your TACO Go code repository and get code into it: `go get`
4. Handle dependencies with the Go Dep package: install Go Dep via `brew install dep` then `brew upgrade dep`.
5. Add and install your dependencies for your Go TACO repository by running `dep init`.

## Running the Go Code locally without a build


```shell
run main.go
```

## Building to TACO Binary

### Building for Docker
```shell
docker build -t taco  .
docker run -p 8080:8080 taco
```

### Build for the local OS
```shell
% go get -t
% go build
```

## Running the TACO Binary

First start up DynamoDB:
```shell
SERVICES=dynamodb localstack start
```

Then create the table:
```shell
awslocal dynamodb create-table --table-name resources \
  --attribute-definitions "AttributeName=id,AttributeType=S" \
  --key-schema "AttributeName=id,KeyType=HASH" \
  --provisioned-throughput=ReadCapacityUnits=100,WriteCapacityUnits=100
```

And add a stub record:
```
awslocal dynamodb put-item --table-name resources --item '{"id": {"S":"99"}, "title":{"S":"Ta-da!"}}'
```

```shell
% AWS_ACCESS_KEY_ID=999999 AWS_SECRET_KEY=1231 ./taco -e development
```

Now visit: http://localhost:8080/v1/resource/99

## Testing

```shell
% go test -v ./...
```
