# Distributed Ordering System using gRPC

This Go-based gRPC project aims to facilitate efficient communication between clients and servers through four distinct methods: unary, server-streaming, client-streaming, and bidirectional streaming for fetching order data. Leveraging the gRPC framework, it enables seamless communication between distributed systems, allowing for real-time data retrieval and processing. Through these different paradigms, the project showcases the versatility and scalability of gRPC in modern software development.

- [Requirements](#requirements)
- [Structure and Implementation](#structure-and-implementation)
- [Results](#results)
  - [Unary](#unary)
  - [Server Stream](#server-stream)
  - [Client Stream](#client-stream)
  - [Bidirectional Stream](#bidirectional-stream)
- [How to run](#how-to-run)

## Requirements

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
