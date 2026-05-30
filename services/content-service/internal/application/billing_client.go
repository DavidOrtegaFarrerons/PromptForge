package application

import "context"

type BillingClient interface {
	ReservePromptSlot(ctx context.Context, userID string) error
	ReleasePromptSlot(ctx context.Context, userID string) error
}
