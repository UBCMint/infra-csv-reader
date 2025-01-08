package server

import (
    "log"
    "net"

    pb "grpc-csv-reader/proto" 
    "google.golang.org/grpc"
)

func StartServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterCSVReaderServiceServer(s, &CSVReaderServer{})

    log.Println("Server is running on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
