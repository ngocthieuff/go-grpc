Guide for arch linux:

error:
```console
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
```

install protoc gen go:
```console
git clone https://aur.archlinux.org/protoc-gen-go.git
ls
makepkg -si
```

error:
```console
--go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
```

install protoc-gen-go-grpc
```console
git clone https://aur.archlinux.org/packages/protoc-gen-go-grpc
ls 
makepkg -si
```

generate model(_grpc.pb.go) and service(pb.go):
```console
mkdir schema
protoc calculator/calculatorpb/calculator.proto --go_out=./schema --go-grpc_out=./schema
```

You can add rule to Makefile and call alias like this:
```console
make gen-cal
```