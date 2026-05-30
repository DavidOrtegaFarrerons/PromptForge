package grpctransport

import (
	"context"
	"errors"

	"github.com/DavidOrtegaFarrerons/promptforge/proto/billing"
	"github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/application"
	"github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BillingGRPCServer struct {
	billing.UnimplementedBillingServer
	reservePromptSlotService application.ReservePromptSlotService
	releasePromptSlotService application.ReleasePromptSlotService
}

func NewBillingGRPCServer(reservePromptSlotService application.ReservePromptSlotService, releasePromptSlotService application.ReleasePromptSlotService) *BillingGRPCServer {
	return &BillingGRPCServer{
		reservePromptSlotService: reservePromptSlotService,
		releasePromptSlotService: releasePromptSlotService,
	}
}

func (s *BillingGRPCServer) ReservePromptSlot(ctx context.Context, req *billing.ReservePromptSlotRequest) (*billing.ReservePromptSlotResponse, error) {
	input := application.ReservePromptSlotInput{UserID: req.UserId}
	err := s.reservePromptSlotService.Execute(ctx, input)
	if err != nil {
		if errors.Is(err, domain.ErrPromptLimitReached) {
			return nil, status.Error(codes.ResourceExhausted, "prompt limit reached")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &billing.ReservePromptSlotResponse{}, nil
}

func (s *BillingGRPCServer) ReleasePromptSlot(ctx context.Context, req *billing.ReleasePromptSlotRequest) (*billing.ReleasePromptSlotResponse, error) {
	input := application.ReleasePromptSlotInput{UserID: req.UserId}
	err := s.releasePromptSlotService.Execute(ctx, input)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &billing.ReleasePromptSlotResponse{}, nil
}
