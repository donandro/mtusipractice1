package handler

import (
	"context"

	pillsservicegrpc "pillsreminder/internal/client/pillsservice/grpc"

	"github.com/rs/zerolog"

	"google.golang.org/grpc"
)

type serverAPI struct {
	pillsservicegrpc.UnimplementedPillsServiceServer
	logger zerolog.Logger
}

func Register(gRPCServer *grpc.Server, logger zerolog.Logger) {
	pillsservicegrpc.RegisterPillsServiceServer(gRPCServer, &serverAPI{
		logger: logger,
	})
}

func (s serverAPI) SearchPills(ctx context.Context, request *pillsservicegrpc.SearchPillsRequest) (*pillsservicegrpc.SearchPillsResponse, error) {
	return &pillsservicegrpc.SearchPillsResponse{
		Pills: []*pillsservicegrpc.Pill{
			{
				Id:          1,
				Name:        "Нурофен-ЭКСТРА",
				Dosage:      "200 мг",
				Instruction: "Инструкция Нурофен-ЭКСТРА",
			},
			{
				Id:          2,
				Name:        "Тирозол",
				Dosage:      "5 мг",
				Instruction: "Инструкция Тирозол",
			},
		},
	}, nil
}

func (s serverAPI) GetPillById(context.Context, *pillsservicegrpc.GetPillByIdRequest) (*pillsservicegrpc.GetPillByIdResponse, error) {
	return &pillsservicegrpc.GetPillByIdResponse{
		Pill: &pillsservicegrpc.Pill{
			Id:          2,
			Name:        "Тирозол",
			Dosage:      "5 мг",
			Instruction: "Инструкция Тирозол",
		},
	}, nil
}
