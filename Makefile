gen-cal:
	protoc calculator/calculatorpb/calculator.proto --go_out=./schema --go-grpc_out=./schema
