# setup local environment to test the package
```
go mod edit -replace omatech.com/urlchecker=../go-urlchecker
go mod tidy
```

# start up and download dependencies
```
go mod init main
go clean -modcache
go get github.com/aws/aws-lambda-go/lambda
go get github.com/aws/aws-sdk-go/aws
go get github.com/aws/aws-sdk-go/aws/session
go get github.com/aws/aws-sdk-go/service/s3
go get github.com/aws/aws-lambda-go/events
go get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
go get github.com/omatech/urlchecker
```

# run with
```
export URL_CHECKER_SECRET="Test Secret String"
go run main.go
```

# test with 
```
export URL_CHECKER_SECRET="Test Secret String"
go test
```

# compile with:

## MacOS
```
GOOS=linux go build -o main main.go
zip main.zip main
```

## Windows
```
$env:GOOS = "linux"
$env:CGO_ENABLED = "0"
$env:GOARCH = "amd64"
go build -o main main.go
```

Then zip only main to main.zip

## Environment variables

Mandatory:
```
URL_CHECKER_SECRET="Super secret String"
```

Optional (default):
```
MAX_SECONDS	(172800)
TIMEZONE	(UTC)
SOURCE_URL_BASE	(https://www.google.com)
TOKEN_MISMATCH_ERROR	(Token mismatch)
URL_TOO_OLD_ERROR	(Url too old)
```




