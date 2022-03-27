/*
 * Copyright © 2022 Durudex

 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.

 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package http

import (
	"errors"
	"strings"

	"github.com/durudex/durudex-gateway/internal/domain"

	"github.com/gofiber/fiber/v2"
)

const authorizationHeader string = "Authorization"

var (
	ErrAuthHeader       = errors.New("invalid auth header")
	ErrAuthTokenIsEmpty = errors.New("token is empty")
)

// HTTP authorization middleware.
func (h *Handler) authMiddleware(ctx *fiber.Ctx) error {
	// Set ip to context value.
	ctx.Context().SetUserValue(domain.IPCtx, ctx.IP())

	// Getting authorization header.
	header := ctx.Get(authorizationHeader)
	if header == "" {
		return ctx.Next()
	}

	// Checking header parts.
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return ErrAuthHeader
	}

	// Check the second part of the header.
	if len(headerParts[1]) == 0 {
		return ErrAuthTokenIsEmpty
	}

	// Parsing jwt access token.
	customClaim, err := h.auth.JWT.Parse(headerParts[1])
	if err != nil {
		return err
	}

	ctx.Context().SetUserValue(domain.UserCtx, customClaim)

	return ctx.Next()
}
