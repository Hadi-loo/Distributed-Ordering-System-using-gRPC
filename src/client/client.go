package main

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "orderingSystem/src/proto"
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
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}(connection)

	client := pb.NewOrderManagementClient(connection)
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
		} else if rpcMode == "unary" {
			fmt.Println("RPC mode is set to unary")
			var orderName string
			fmt.Printf("Enter order name: ")
			fmt.Scanln(&orderName)
			newOrder := &pb.OrderRequest{OrderID: index, OrderName: orderName}
			index += 1
			res, err := client.UnaryGetOrder(ctx, newOrder)
			if err != nil {
				log.Printf("Failed to call UnaryGetOrder: %v\n", err)
			}
			fmt.Printf("Received: %d, %s, %s\n", res.GetOrderId(), res.GetOrderName(), res.GetOrderTimestamp())
		} else if rpcMode == "server_stream" {
			fmt.Println("RPC mode is set to server streaming")
			var orderName string
			fmt.Printf("Enter order name: ")
			fmt.Scanln(&orderName)
			newOrder := &pb.OrderRequest{OrderID: index, OrderName: orderName}
			index += 1
			serverStream, err := client.ServerStreamGetOrder(ctx, newOrder)
			if err != nil {
				log.Printf("Failed to call ServerStreamGetOrder: %v\n", err)
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
		} else if rpcMode == "client_stream" {
			fmt.Println("RPC mode is set to client streaming")
			clientStream, err := client.ClientStreamGetOrder(ctx)
			if err != nil {
				log.Printf("Failed to call ClientStreamGetOrder: %v\n", err)
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
				if err := clientStream.Send(&pb.OrderRequest{OrderID: index, OrderName: order}); err != nil {
					log.Printf("Failed to send order: %v\n", err)
				}
				index += 1
			}
			res, err := clientStream.CloseAndRecv()
			if err != nil {
				log.Printf("Failed to receive order: %v\n", err)
			}
			fmt.Printf("Received: %d, %s, %s\n", res.GetOrderId(), res.GetOrderName(), res.GetOrderTimestamp())
		} else if rpcMode == "bidi_stream" {
			fmt.Println("RPC mode is set to bi-directional streaming")
			bidiStream, err := client.BiDiStreamGetOrder(ctx)
			if err != nil {
				log.Printf("Failed to call BiDiStreamGetOrder: %v\n", err)
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
				if err := bidiStream.Send(&pb.OrderRequest{OrderID: index, OrderName: orderName}); err != nil {
					log.Printf("Failed to send order: %v\n", err)
					break
				}
			}
		} else {
			fmt.Println("Invalid rpc mode")
		}
	}

	//// unary
	//fmt.Println("Unary RPC")
	//newOrder := &pb.OrderRequest{OrderID: index, OrderName: "apple"}
	//index += 1
	//res, err := client.UnaryGetOrder(ctx, newOrder)
	//if err != nil {
	//	log.Printf("Failed to call UnaryGetOrder: %v\n", err)
	//}
	//fmt.Printf("Received: %d, %s, %s\n", res.GetOrderId(), res.GetOrderName(), res.GetOrderTimestamp())
	//fmt.Println("Unary RPC completed")
	//fmt.Println("========================================")
	//
	//// server stream
	//fmt.Println("Server Streaming RPC")
	//newOrder = &pb.OrderRequest{OrderID: index, OrderName: "apple"}
	//index += 1
	//serverStream, err := client.ServerStreamGetOrder(ctx, newOrder)
	//if err != nil {
	//	log.Printf("Failed to call ServerStreamGetOrder: %v\n", err)
	//}
	//for {
	//	res, err := serverStream.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		log.Printf("Failed to receive order: %v\n", err)
	//	}
	//	fmt.Printf("Received: %d, %s, %s\n", res.GetOrderId(), res.GetOrderName(), res.GetOrderTimestamp())
	//}
	//fmt.Println("Server Streaming RPC completed")
	//fmt.Println("========================================")
	//
	//// client stream
	//fmt.Println("Client Streaming RPC")
	//clientStream, err := client.ClientStreamGetOrder(ctx)
	//if err != nil {
	//	log.Printf("Failed to call ClientStreamGetOrder: %v\n", err)
	//}
	//orders := []*pb.OrderRequest{
	//	{OrderID: index, OrderName: "banana"},
	//	{OrderID: index + 1, OrderName: "apple"},
	//	{OrderID: index + 2, OrderName: "orange"},
	//	{OrderID: index + 3, OrderName: "grape"},
	//}
	//index += 4
	//for _, order := range orders {
	//	if err := clientStream.Send(order); err != nil {
	//		log.Printf("Failed to send order: %v\n", err)
	//	}
	//}
	//res, err = clientStream.CloseAndRecv()
	//if err != nil {
	//	log.Printf("Failed to receive order: %v\n", err)
	//}
	//fmt.Printf("Received: %d, %s, %s\n", res.GetOrderId(), res.GetOrderName(), res.GetOrderTimestamp())
	//fmt.Println("Client Streaming RPC completed")
	//fmt.Println("========================================")
	//
	//// bidirectional stream
	//fmt.Println("Bi-Directional Streaming RPC")
	//bidiStream, err := client.BiDiStreamGetOrder(ctx)
	//waitc := make(chan struct{})
	//go func() {
	//	for {
	//		res, err := bidiStream.Recv()
	//		if err == io.EOF {
	//			close(waitc)
	//			break
	//		}
	//		if err != nil {
	//			log.Printf("Failed to receive order: %v\n", err)
	//		}
	//		fmt.Printf("Received: %d, %s, %s\n", res.GetOrderId(), res.GetOrderName(), res.GetOrderTimestamp())
	//	}
	//}()
	//orders = []*pb.OrderRequest{
	//	{OrderID: index, OrderName: "banana"},
	//	{OrderID: index + 1, OrderName: "apple"},
	//	{OrderID: index + 2, OrderName: "orange"},
	//	{OrderID: index + 3, OrderName: "grape"},
	//}
	//index += 4
	//for _, order := range orders {
	//	if err := bidiStream.Send(order); err != nil {
	//		log.Printf("Failed to send order: %v\n", err)
	//	}
	//}
	//if err := bidiStream.CloseSend(); err != nil {
	//	log.Printf("Failed to close send: %v\n", err)
	//}
	//<-waitc
}
