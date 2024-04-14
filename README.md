# Distributed Ordering System using gRPC

This Go-based gRPC project aims to facilitate efficient communication between client and server through four distinct methods: unary, server-streaming, client-streaming, and bidirectional streaming for fetching order data.Through these different paradigms, the project showcases the versatility and scalability of gRPC in modern software development.

- [Requirements](#requirements)
- [Structure and Implementation](#structure-and-implementation)
- [Results](#results)
  - [Unary](#unary)
  - [Server Stream](#server-stream)
  - [Client Stream](#client-stream)
  - [Bidirectional Stream](#bidirectional-stream)
- [How to run](#how-to-run)

## Requirements

For this project, the installation of the Go programming language is required. Additionally, the dependencies for gRPC packages are specified in the go.mod file, as illustrated in the following box:

```go
module orderingSystem

go 1.22

require (
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
)

require (
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240412170617-26222e5d3d56 // indirect
)

```
You can use the following command to download the requirments metioned in above box:

```go
go mod download
```

## Structure and Implementation

## Results

To obtain results for all four methods, the server and client files are executed, after which requests are dispatched for each method accordingly:

### Unary

#### Client Terminal
![image](https://github.com/Hadi-loo/Distributed-Ordering-System-using-gRPC/assets/88041997/d230c2a1-1776-412b-a9a5-6aa6ad1c2a42)

#### Server Terminal
![image](https://github.com/Hadi-loo/Distributed-Ordering-System-using-gRPC/assets/88041997/9c23ae55-5b73-425f-ab1c-c0206b4a7825)

### Server Stream

#### Client Terminal
![image](https://github.com/Hadi-loo/Distributed-Ordering-System-using-gRPC/assets/88041997/8ee2e64e-d2ae-4df2-8f45-c44a8132a147)

#### Server Terminal
![image](https://github.com/Hadi-loo/Distributed-Ordering-System-using-gRPC/assets/88041997/718f94da-849a-4f1e-a60d-f6626951ebea)

### Client Stream

#### Client Terminal
![image](https://github.com/Hadi-loo/Distributed-Ordering-System-using-gRPC/assets/88041997/f9818a04-7ad4-46c5-880d-86b094b8604d)

#### Server Terminal
![image](https://github.com/Hadi-loo/Distributed-Ordering-System-using-gRPC/assets/88041997/67d1b58c-4503-4477-b60d-6be06346e5ed)

### Bidirectional Stream

#### Client Terminal
![image](https://github.com/Hadi-loo/Distributed-Ordering-System-using-gRPC/assets/88041997/6ae31e69-59f4-4235-8b2a-72687da265df)

#### Server Terminal
![image](https://github.com/Hadi-loo/Distributed-Ordering-System-using-gRPC/assets/88041997/b98b94c2-2f63-443e-89be-f832329d6a80)



## How to run

To run this project, follow these steps:

### 1. Compile Dependencies

Ensure that you have Go installed on your system. Navigate to the root directory of the project where the `go.mod` file is located, and compile the dependencies by executing:

```bash
go mod download
```

### 2. Compile Protocol Buffers

Compile the protocol buffers defined in the `.proto` file using Protocol Buffers compiler (`protoc`). Run the following command:

```bash
protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative src/proto/orderingSystem.proto
```

This command generates Go files from the protocol buffers.(orderingSystem_grpc.pb.go and orderingSystem.pb.go)

### 3. Run Server

To start the server, run the following command:

```bash
go run server/server.go
```

The server will start listening for incoming gRPC requests.

### 4. Run Client

To execute the client code, run the following command:

```bash
go run client/client.go
```

This will send requests to the server and display the responses.
