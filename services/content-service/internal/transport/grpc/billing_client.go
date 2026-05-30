package grpctransport

import (
	"context"

	"github.com/DavidOrtegaFarrerons/promptforge/proto/billing"
	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCBillingClient struct {
	client billing.BillingClient
}

func NewGRPCBillingClient(conn *grpc.ClientConn) *GRPCBillingClient {
	return &GRPCBillingClient{
		client: billing.NewBillingClient(conn),
	}
}

func (c *GRPCBillingClient) ReservePromptSlot(ctx context.Context, userID string) error {
	reservePromptSlotRequest := &billing.ReservePromptSlotRequest{UserId: userID}
	_, err := c.client.ReservePromptSlot(ctx, reservePromptSlotRequest)
	if err != nil {
		if status.Code(err) == codes.ResourceExhausted {
			return domain.ErrPromptLimitReached
		}
		return err
	}

	return nil
}

func (c *GRPCBillingClient) ReleasePromptSlot(ctx context.Context, userID string) error {
	releasePromptSlotRequest := &billing.ReleasePromptSlotRequest{UserId: userID}
	_, err := c.client.ReleasePromptSlot(ctx, releasePromptSlotRequest)

	return err
}
