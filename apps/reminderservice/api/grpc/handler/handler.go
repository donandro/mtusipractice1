package handler

import (
	"context"

	reminderservicegrpc "pillsreminder/internal/client/reminderservice/grpc"

	"github.com/rs/zerolog"

	"google.golang.org/grpc"
)

type serverAPI struct {
	reminderservicegrpc.UnimplementedReminderServiceServer
	logger zerolog.Logger
}

func Register(gRPCServer *grpc.Server, logger zerolog.Logger) {
	reminderservicegrpc.RegisterReminderServiceServer(gRPCServer, &serverAPI{
		logger: logger,
	})
}

func (s serverAPI) AddPlan(ctx context.Context, request *reminderservicegrpc.AddPlanRequest) (*reminderservicegrpc.AddPlanResponse, error) {
	return &reminderservicegrpc.AddPlanResponse{}, nil
}
func (s serverAPI) GetUserPlans(ctx context.Context, request *reminderservicegrpc.GetUserPlansRequest) (*reminderservicegrpc.GetUserPlansResponse, error) {
	return &reminderservicegrpc.GetUserPlansResponse{
		UserPlans: []*reminderservicegrpc.UserPlan{
			{
				PlanId:       1,
				Name:         "Первый план",
				PillId:       2,
				ReminderTime: "10:00:00.000 GMT+3",
				Frequency:    "daily",
				DaysOfWeek:   "Mon,Wed,Fri",
				StartDate:    "2024-09-01",
				EndDate:      "2024-09-30",
			},
		},
	}, nil
}
func (s serverAPI) DeleteUserPlan(ctx context.Context, request *reminderservicegrpc.DeleteUserPlanRequest) (*reminderservicegrpc.DeleteUserPlanResponse, error) {
	return &reminderservicegrpc.DeleteUserPlanResponse{}, nil
}
func (s serverAPI) CancelPlan(ctx context.Context, request *reminderservicegrpc.CancelPlanRequest) (*reminderservicegrpc.CancelPlanResponse, error) {
	return &reminderservicegrpc.CancelPlanResponse{}, nil
}
func (s serverAPI) Intake(ctx context.Context, request *reminderservicegrpc.IntakeRequest) (*reminderservicegrpc.IntakeResponse, error) {
	return &reminderservicegrpc.IntakeResponse{}, nil
}
