package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"
)

func (r *mutationResolver) SignUp(ctx context.Context, input domain.SignUpInput) (uint64, error) {
	return r.service.Auth.SignUp(ctx, input)
}

func (r *mutationResolver) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	return r.service.Auth.SignIn(ctx, input)
}

func (r *mutationResolver) RefreshTokens(ctx context.Context, input domain.RefreshTokensInput) (*domain.Tokens, error) {
	return r.service.Auth.RefreshTokens(ctx, input)
}
