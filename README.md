# Distributed Ordering System using gRPC

This Go-based gRPC project aims to facilitate efficient communication between client and server through four distinct methods: unary, server-streaming, client-streaming, and bidirectional streaming for fetching order data.Through these different paradigms, the project showcases the versatility and scalability of gRPC in modern software development.

- [Requirements](#requirements)
- [Structure and Implementation](#structure-and-implementation)
  - [.proto file](#proto-file)
  - [Client](#client)
  - [Server](#server)
    - [Server Initialization](#server-initialization)
    - [RPC Methods](#rpc-methods)
    - [SearchItems Function](#searchitems-function)
    - [Database](#database)
- [Results](#results)
  - [Unary RPC](#unary-rpc)
  - [Server Streaming RPC](#server-streaming-rpc)
  - [Client Streaming RPC](#client-streaming-rpc)
  - [Bidirectional Streaming RPC](#bidirectional-streaming-rpc)
- [Comparison and Usage](#comparison-and-usage)
  - [Unary](#unary)
  - [Server Stream](#server-stream)
  - [Client Stream](#client-stream)
  - [Bidirectional Stream](#bidirectional-stream)
- [How to run](#how-to-run)
  - [1. Compile Dependencies](#1-compile-dependencies)
  - [2. Compile Protocol Buffers](#2-compile-protocol-buffers)
  - [3. Run Server](#3-run-server)
  - [4. Run Client](#4-run-client)
- [Contributions](#contributions)

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

### .proto file

The `.proto` file is written in proto3 syntax, which is a language-neutral, platform-neutral, extensible way of serializing structured data for use in communications protocols, data storage, and more. Here's some explanation on the .proto file code and its functionality:

```proto
service OrderManagement {
  rpc UnaryGetOrder(OrderRequest) returns (OrderResponse) {}
  rpc ServerStreamGetOrder(OrderRequest) returns (stream OrderResponse) {}
  rpc ClientStreamGetOrder(stream OrderRequest) returns (OrderResponse) {}
  rpc BiDiStreamGetOrder(stream OrderRequest) returns (stream OrderResponse) {}
}
```

We defined a service with RPC (Remote Procedure Call) methods. Inside the `OrderManagement` service, there are four different RPC methods defined:

1. `UnaryGetOrder:` A simple RPC where the client sends a single OrderRequest and gets a single OrderResponse.

2. `ServerStreamGetOrder:` A server streaming RPC where the client sends a single OrderRequest and gets a stream of OrderResponse messages back.

3. `ClientStreamGetOrder:` A client streaming RPC where the client sends a stream of OrderRequest messages and gets a single OrderResponse.

4. `BiDiStreamGetOrder:` A bidirectional streaming RPC where both the client and server send a stream of messages to each other.

```proto
message OrderRequest {
  int32 OrderID = 1;
  string OrderName = 2;
}
```
Here, we defined a message type. OrderRequest is a message containing an `OrderID` of type int32 and an `OrderName` of type string. Both of them are labeled with unique tags (1 and 2, respectively).

The numbers assigned to each field (e.g., = 1, = 2) are field tags used to identify your fields in the message binary format and should be unique within a message type. They are essential for the backward compatibility of your message type.

```proto
message OrderResponse {
  int32 OrderId = 1;
  string OrderName = 2;
  string OrderTimestamp = 3;
}
```
Similarly, OrderResponse is a message type that includes an `OrderId`, `OrderName`, and an `OrderTimestamp`, all of which are labeled with unique tags (1, 2, and 3 respectively). These tags are used in the binary encoding of the message and should not be changed once your message type is in use.

Overall, this `.proto` file defines the structure of messages and services for an ordering system that can handle different types of communication patterns between a client and a server. The generated code from this `.proto` file will be used by the client and server to serialize, send, and receive the defined messages.

### Client

The `client.go` file contains the code for a client application that communicates with a server using gRPC, a high-performance, open-source universal RPC framework. Here's a breakdown of the code and its functions:

First, we defined constants for the server address, port, and a timeout value:

```go
const (
	address          = "localhost"
	port             = ":50505"
	timeOutInSeconds = 1200
)
```

The `main()` function is the entry point of the program. It establishes a connection to the gRPC server, creates a client from the connection, and processes user input for different RPC modes.
```go
func main() { ... }
```
The main function also includes a loop that prompts the user to enter the desired RPC mode (unaryMode, serverStreamMode, clientStreamMode, bidiStreamMode) and calls the corresponding function to handle the request. The loop continues until the user enters exit. Based on the user's input, the program switches between different RPC modes. If an invalid mode is entered, it prints an error message.

The actual RPC calls would be handled in the functions that are called based on the user's choice of RPC mode:

1. **Unary Mode**:
```go
func unaryMode() { ... }
```
This function handles the unary RPC mode where a single request is sent to the server and a single response is received.

2. **Server Stream Mode**:
```go
func serverStreamMode() { ... }
```
In this function, we handled the server streaming RPC mode where a single request is sent to the server and a stream of responses is received.

3. **Client Stream Mode**:
```go
func clientStreamMode() { ... }
```
Client streaming RPC mode where a stream of requests is sent to the server and a single response is received, is handled in this function.

4. **Bidirectional Stream Mode**:
```go
func bidiStreamMode() { ... }
```
This function handles the bidirectional streaming RPC mode where both the client and server can send a stream of messages to each other.

### Server

`server.go` file contains the implementation of a gRPC server for an ordering system. It provides various RPC methods to handle client requests. Now let's delve into important parts of this code: 

#### Server Initialization
The main function initializes the server by:
- Starting the gRPC server.
- Registering the OrderManagement service with the server.
- Listening for incoming connections(The server listens on port 50505 for incoming gRPC requests.).

#### RPC Methods
1. **UnaryGetOrder**: 
   - Handles unary RPC requests.
   - Searches for items based on the order name provided by the client.
   - Returns the first found item along with its details.

2. **ServerStreamGetOrder**: 
   - Handles server streaming RPC requests.
   - Searches for items based on the order name provided by the client.
   - Streams back all found items to the client.

3. **ClientStreamGetOrder**: 
   - Handles client streaming RPC requests.
   - Receives multiple order names from the client.
   - Searches for each item and streams back the first found item for each order name.

4. **BiDiStreamGetOrder**: 
   - Handles bidirectional streaming RPC requests.
   - Receives order names from the client.
   - Streams back all found items for each order name.

#### SearchItems Function
- Helper function to search for items based on a search string.
- Used by various RPC methods to search for items in the database.

#### Database
The database package (`orderingSystem/database`) contains a predefined list of items.

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


## Comparison and Usage

### Unary RPC

- Simplest form of RPC (Remote Procedure Call) in gRPC, where the client sends a single request to the server and gets back a single response.
- Suitable for scenarios where you need to retrieve or manipulate a single piece of data with a single request-response interaction. For example, fetching a user's profile, retrieving a specific record from a 
  database, or performing a simple calculation.

### Server Streaming RPC

- Server streaming RPC allows the server to send multiple responses to a single client request. The client initiates the request, and the server responds with a stream of messages.
- Useful when the client needs to receive a potentially large amount of data that the server can generate in a sequential manner. For instance, fetching a feed of updates, processing a large file 
  in chunks, or real-time data feeds like stock prices.

### Client Streaming RPC

- Enables the client to send a stream of messages to the server, and then the server responds with a single message (once it has processed the entire stream).
- Appropriate when the client needs to send a significant amount of data to the server for processing, such as uploading a large file, sending telemetry data from IoT devices, or sending logs or 
  metrics from multiple sources.

### Bidirectional Streaming RPC

-  Allows both the client and the server to send a stream of messages to each other. This means that the client can send messages to the server at any time, and the server can also send messages to the client at any time.
- Beneficial in scenarios where there is a need for real-time communication and both client and server need to send and receive data concurrently. Examples include chat applications, multiplayer games, or collaborative editing applications.


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


## Contributions

- **Server File**: Hadi Babaloo, Sina Tabasi
- **Client File**: Hadi Babaloo, Kasra Haji-Heydari
- **proto File**: Hadi Babaloo, MohammadSadegh Aboofazeli
- **Report and Analysis**: MohammadSadegh Aboofazeli, Sina Tabasi, Kasra Haji-Heydari
