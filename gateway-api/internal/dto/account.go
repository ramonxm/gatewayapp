package dto

import (
	"time"

	"github.com/ramonxm/gatewayapp/gateway-api/internal/domain"
)

type CreateAccountInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ApiKey    string    `json:"api_key,omitempty"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToAccount(input CreateAccountInput) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

func FromAccount(account *domain.Account) *AccountOutput {
	return &AccountOutput{
		ID:        account.Id,
		Name:      account.Name,
		Email:     account.Email,
		ApiKey:    account.APIKey,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
