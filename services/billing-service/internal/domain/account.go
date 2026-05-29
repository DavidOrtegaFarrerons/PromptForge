package domain

import "time"

type AccountID string
type UserID string

type Plan string

const (
	PlanFree Plan = "free"
	PlanPro  Plan = "pro"
)

type Account struct {
	accountID AccountID
	userID    UserID
	plan      Plan
	createdAt time.Time
	updatedAt time.Time
}

func (a Account) AccountID() AccountID {
	return a.accountID
}

func (a Account) UserID() UserID {
	return a.userID
}

func (a Account) Plan() Plan {
	return a.plan
}

func (a Account) CreatedAt() time.Time {
	return a.createdAt
}

func (a Account) UpdatedAt() time.Time {
	return a.updatedAt
}

func NewAccount(accountID AccountID, userID UserID, plan Plan, createdAt time.Time, updatedAt time.Time) (Account, error) {
	if accountID == "" {
		return Account{}, ErrEmptyAccount
	}

	if userID == "" {
		return Account{}, ErrEmptyUserID
	}

	if !isValidPlan(plan) {
		return Account{}, ErrPlanInvalid
	}

	return Account{
		accountID: accountID,
		userID:    userID,
		plan:      plan,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil

}

func isValidPlan(p Plan) bool {
	return p == PlanPro || p == PlanFree
}
