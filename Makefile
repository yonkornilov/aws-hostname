export GOOS=linux

BIN=aws-hostname

$(BIN): main.go
	go build -o $(BIN) main.go

test:
	AWS_REGION="us-east-1" go test
