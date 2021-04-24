# grpc server

chat style app to learn the basics of grpc.
## Installation

Please make sure you have the protoc buffer installed in your $GOPATH
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Usage

in split terminals run
```bash
go run main.go
```
on one and 
```bash
go run client.go
```
to display the connection message

## What I learned:
* Protocol buffers
* How to write services
* GOPATH
* Messages
* Context and Googles GRPC packages

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
