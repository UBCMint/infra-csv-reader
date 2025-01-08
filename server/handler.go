package server

import (
	"os"

	pb "grpc-csv-reader/proto" 
	"grpc-csv-reader/utils"  
)

type CSVReaderServer struct {
    pb.UnimplementedCSVReaderServiceServer
}

func (s *CSVReaderServer) GetCSVData(req *pb.Empty, stream pb.CSVReaderService_GetCSVDataServer) error {
    file, err := os.Open("csv/EEG_recording_2024-03-23-04.33.02.csv")
    if err != nil {
        return err
    }
    defer file.Close()

    rows, err := utils.ReadCSV(file)
    if err != nil {
        return err
    }

    for _, row := range rows {
        if err := stream.Send(&pb.CSVRow{
            Column1: row[0],
            Column2: row[1],
            Column3: row[2],
        }); err != nil {
            return err
        }
    }
    return nil
}
