/*
	Copyright © 2021-2022 Durudex

	This file is part of Durudex: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	Durudex is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with Durudex. If not, see <https://www.gnu.org/licenses/>.
*/

package service

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/delivery/grpc"
	"github.com/durudex/durudex-gateway/internal/domain"
)

type Auth interface {
	SignUp(ctx context.Context, input *domain.SignUpInput) (uint64, error)
	SignIn(ctx context.Context, input *domain.SignInInput) (domain.Tokens, error)
	Verify(ctx context.Context, input *domain.VerifyInput) (bool, error)
	GetCode(ctx context.Context, code uint64) (bool, error)
	RefreshTokens(ctx context.Context, input *domain.RefreshTokensInput) (domain.Tokens, error)
}

type Service struct {
	Auth
}

// Creating a new service.
func NewService(grpcHandler *grpc.Handler) *Service {
	return &Service{Auth: NewAuthService(grpcHandler)}
}
