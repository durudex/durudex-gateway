/*
	Copyright © 2021 Durudex

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

package graphql

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	authorizationHeader = "Authorization"

	// Ctx claims.
	userCtx = "userID"
)

// Middleware validation of authorization header for validity.
func (h *Handler) authMiddleware(c *fiber.Ctx) error {
	// Getting header.
	header := c.Get(authorizationHeader)
	if header == "" {
		return c.Next()
	}

	// Divide the header into two parts.
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return errors.New("invalid auth header")
	}

	// Check the second part of the header.
	if len(headerParts[1]) == 0 {
		return errors.New("token is empty")
	}

	// Parsing header.
	customClaim, err := h.auth.JWT.Parse(headerParts[1])
	if err != nil {
		return err
	}

	// Set claims value.
	c.Context().SetUserValue(userCtx, customClaim)

	return c.Next()
}

// User identification vy ctx.
func (h *Handler) userIdentity(ctx context.Context) bool {
	// Checking ctx value on nil.
	if ctx.Value(userCtx) == nil {
		return false
	}

	return true
}

// Get user id by ctx.
func GetUserID(ctx context.Context) (uint64, error) {
	return strconv.ParseUint(ctx.Value(userCtx).(string), 10, 64)
}
