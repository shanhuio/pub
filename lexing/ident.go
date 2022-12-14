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

// IsIdentLetter checks if rune r can start an identifier.
func IsIdentLetter(r rune) bool {
	return r == '_' || IsLetter(r)
}

// LexIdent lexes a typical C/Go langauge identifier.
func LexIdent(x *Lexer, t int) *Token {
	if !IsIdentLetter(x.Rune()) {
		panic("ident must start with letter or _")
	}

	for {
		r, _ := x.Next()
		if !IsIdentLetter(r) && !IsDigit(r) {
			break
		}
	}
	return x.MakeToken(t)
}
