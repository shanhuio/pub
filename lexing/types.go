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

package lexing

import (
	"fmt"
)

// Types is a registrar of token type names
type Types struct {
	names map[int]string
}

// NewTypes makes a new registrar. It will auto register the default tokens.
func NewTypes() *Types {
	ret := new(Types)
	ret.names = make(map[int]string)

	ret.Register(EOF, "eof")
	ret.Register(Comment, "comment")
	ret.Register(Illegal, "illegal")

	return ret
}

// Register registers a type with a name.
// If the type is already registered, it panics.
func (types *Types) Register(t int, name string) {
	if _, found := types.names[t]; found {
		panic("token already registered")
	}

	types.names[t] = name
}

// Name resolves the name of a type.
func (types *Types) Name(t int) string {
	if ret, found := types.names[t]; found {
		return ret
	}

	return fmt.Sprintf("<T%d>", t)
}
