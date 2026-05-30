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
	accountID   AccountID
	userID      UserID
	plan        Plan
	promptCount int
	createdAt   time.Time
	updatedAt   time.Time
}

func NewAccount(accountID AccountID, userID UserID, plan Plan, promptCount int, createdAt time.Time, updatedAt time.Time) (Account, error) {
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
		accountID:   accountID,
		userID:      userID,
		plan:        plan,
		promptCount: promptCount,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}, nil

}

func isValidPlan(p Plan) bool {
	return p == PlanPro || p == PlanFree
}

func (a Account) CanCreatePrompt() bool {
	switch a.plan {
	case PlanFree:
		return a.promptCount < 10
	case PlanPro:
		return true
	default:
		return false
	}
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

func (a Account) PromptCount() int {
	return a.promptCount
}

func (a Account) CreatedAt() time.Time {
	return a.createdAt
}

func (a Account) UpdatedAt() time.Time {
	return a.updatedAt
}
