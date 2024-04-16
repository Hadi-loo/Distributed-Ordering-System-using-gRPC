package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"orderingSystem/src/proto"
	"time"

	"google.golang.org/grpc"
)

const (
	address          = "localhost"
	port             = ":50505"
	timeOutInSeconds = 1200
)

func main() {
	connection, err := grpc.Dial(address+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer func() {
		if err := connection.Close(); err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}()

	client := proto.NewOrderManagementClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), timeOutInSeconds*time.Second)
	defer cancel()

	var index int32 = 1

	fmt.Printf("You can change the rpcMode variable using the -mode flag\n")
	fmt.Printf("Possible values: unary, server_stream, client_stream, bidi_stream\n")

	for {
		var rpcMode string
		fmt.Printf("\nEnter rpc mode: ")
		fmt.Scanln(&rpcMode)
		if rpcMode == "exit" {
			break
		}

		switch rpcMode {
		case "unary":
			unaryMode(client, ctx, &index)
		case "server_stream":
			serverStreamMode(client, ctx, &index)
		case "client_stream":
			clientStreamMode(client, ctx, &index)
		case "bidi_stream":
			bidiStreamMode(client, ctx, &index)
		default:
			fmt.Println("Invalid rpc mode")
		}
	}
}

func unaryMode(client proto.OrderManagementClient, ctx context.Context, index *int32) {
	fmt.Println("RPC mode is set to unary")
	var orderName string
	fmt.Printf("Enter order name: ")
	fmt.Scanln(&orderName)
	newOrder := &proto.OrderRequest{OrderID: *index, OrderName: orderName}
	*index++
	res, err := client.UnaryGetOrder(ctx, newOrder)
	if err != nil {
		log.Printf("Failed to call UnaryGetOrder: %v\n", err)
		return
	}
	fmt.Printf("Received: %d, %s, %s\n", res.GetOrderId(), res.GetOrderName(), res.GetOrderTimestamp())
}

func serverStreamMode(client proto.OrderManagementClient, ctx context.Context, index *int32) {
	fmt.Println("RPC mode is set to server streaming")
	var orderName string
	fmt.Printf("Enter order name: ")
	fmt.Scanln(&orderName)
	newOrder := &proto.OrderRequest{OrderID: *index, OrderName: orderName}
	*index++
	serverStream, err := client.ServerStreamGetOrder(ctx, newOrder)
	if err != nil {
		log.Printf("Failed to call ServerStreamGetOrder: %v\n", err)
		return
	}
	for {
		res, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Failed to receive order: %v\n", err)
			break
		}
		fmt.Printf("Received: %d, %s, %s\n", res.GetOrderId(), res.GetOrderName(), res.GetOrderTimestamp())
	}
}

func clientStreamMode(client proto.OrderManagementClient, ctx context.Context, index *int32) {
	fmt.Println("RPC mode is set to client streaming")
	clientStream, err := client.ClientStreamGetOrder(ctx)
	if err != nil {
		log.Printf("Failed to call ClientStreamGetOrder: %v\n", err)
		return
	}
	var orders []string
	for {
		var orderName string
		fmt.Printf("Enter order name: ")
		fmt.Scanln(&orderName)
		if orderName == "exit" {
			break
		}
		orders = append(orders, orderName)
	}
	for _, order := range orders {
		if err := clientStream.Send(&proto.OrderRequest{OrderID: *index, OrderName: order}); err != nil {
			log.Printf("Failed to send order: %v\n", err)
			break
		}
		*index++
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Printf("Failed to receive order: %v\n", err)
		return
	}
	fmt.Printf("Received: %d, %s, %s\n", res.GetOrderId(), res.GetOrderName(), res.GetOrderTimestamp())
}

func bidiStreamMode(client proto.OrderManagementClient, ctx context.Context, index *int32) {
	fmt.Println("RPC mode is set to bi-directional streaming")
	bidiStream, err := client.BiDiStreamGetOrder(ctx)
	if err != nil {
		log.Printf("Failed to call BiDiStreamGetOrder: %v\n", err)
		return
	}
	waitc := make(chan struct{})
	go func() {
		for {
			res, err := bidiStream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Printf("Failed to receive order: %v\n", err)
				// TODO: close the channel
				close(waitc)
				return
			}
			fmt.Printf("Received: %d, %s, %s\n", res.GetOrderId(), res.GetOrderName(), res.GetOrderTimestamp())
		}
	}()
	for {
		var orderName string
		fmt.Scanln(&orderName)
		if orderName == "exit" {
			break
		}
		if err := bidiStream.Send(&proto.OrderRequest{OrderID: *index, OrderName: orderName}); err != nil {
			log.Printf("Failed to send order: %v\n", err)
			break
		}
		*index++
	}
}
