// Copyright (C) 2023  Shanhu Tech Inc.
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
// for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package oauth2

import (
	"shanhu.io/pub/aries"
)

// UserMeta returns the meta data returned by a sign in.
type UserMeta struct {
	Method string
	ID     string
	Name   string // Screen name.
	Email  string
}

type metaExchange interface {
	callback(c *aries.C) (*UserMeta, *State, error)
}

type provider interface {
	metaExchange
	client() *Client
}
