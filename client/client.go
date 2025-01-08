package main

import (
    "context"
    "io"
    "log"

    pb "grpc-csv-reader/proto"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect to server: %v", err)
    }
    defer conn.Close()

    client := pb.NewCSVReaderServiceClient(conn)

    log.Println("Requesting CSV data...")
    stream, err := client.GetCSVData(context.Background(), &pb.Empty{})
    if err != nil {
        log.Fatalf("Error calling GetCSVData: %v", err)
    }

    for {
        row, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("Error receiving stream: %v", err)
        }

        log.Printf("Received row: %s, %s, %s", row.GetColumn1(), row.GetColumn2(), row.GetColumn3())
    }
}
