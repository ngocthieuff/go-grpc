gen-cal:
	protoc calculator/calculatorpb/calculator.proto --go_out=./schema --go-grpc_out=./schema
run-server:
	go run server.go
