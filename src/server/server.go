package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	db "orderingSystem/database"
	pb "orderingSystem/src/proto"
	"strings"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":50505"
)

type server struct {
	pb.UnimplementedOrderManagementServer
}

func SearchItems(searchString string) []string {
	var foundItems []string
	for _, item := range db.Items {
		if strings.Contains(item, searchString) {
			foundItems = append(foundItems, item)
		}
	}
	return foundItems
}

func (s *server) UnaryGetOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	fmt.Printf("Received order for %v\n", req.GetOrderName())
	var foundItems = SearchItems(req.GetOrderName())
	if len(foundItems) == 0 {
		return nil, fmt.Errorf("item not found")
	}
	return &pb.OrderResponse{OrderId: req.GetOrderID(), OrderName: foundItems[0], OrderTimestamp: time.Now().String()}, nil
}

func (s *server) ServerStreamGetOrder(req *pb.OrderRequest, stream pb.OrderManagement_ServerStreamGetOrderServer) error {
	fmt.Printf("Received order for %v\n", req.GetOrderName())
	var foundItems = SearchItems(req.GetOrderName())
	if len(foundItems) == 0 {
		return fmt.Errorf("item not found")
	}
	for _, item := range foundItems {
		res := &pb.OrderResponse{OrderId: req.GetOrderID(), OrderName: item, OrderTimestamp: time.Now().String()}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) ClientStreamGetOrder(stream pb.OrderManagement_ClientStreamGetOrderServer) error {
	var orders []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Received order for %v\n", req.GetOrderName())
		orders = append(orders, req.GetOrderName())
	}
	var err error
	for _, order := range orders {
		var foundItems = SearchItems(order)
		if len(foundItems) == 0 {
			err = fmt.Errorf("item not found")
		} else {
			for _, item := range foundItems {
				res := &pb.OrderResponse{OrderId: 1, OrderName: item, OrderTimestamp: time.Now().String()}
				err = stream.SendAndClose(res)
				if err == nil {
					return nil
				}
			}
		}
	}
	return err
}

func (s *server) BiDiStreamGetOrder(stream pb.OrderManagement_BiDiStreamGetOrderServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("Received order for %v\n", req.GetOrderName())
		var foundItems = SearchItems(req.GetOrderName())
		if len(foundItems) == 0 {
			return fmt.Errorf("item not found")
		}
		for _, item := range foundItems {
			res := &pb.OrderResponse{OrderId: req.GetOrderID(), OrderName: item, OrderTimestamp: time.Now().String()}
			if err := stream.Send(res); err != nil {
				return err
			}
		}
	}
}

func main() {

	fmt.Println("Starting server...")

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register the service with the server
	pb.RegisterOrderManagementServer(s, &server{})

	// Start the server
	if err := s.Serve(listener); err != nil {
		fmt.Println("Failed to serve:", err)
	}

	fmt.Println("Server started!")

}
