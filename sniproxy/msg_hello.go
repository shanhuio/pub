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

package sniproxy

type helloRequest struct {
	msg string
}

func (m *helloRequest) encodeTo(enc *encoder) {
	enc.str(m.msg)
}

func (m *helloRequest) decodeFrom(dec *decoder) {
	m.msg = dec.str()
}

type helloResponse struct {
	msg string
}

func (m *helloResponse) encodeTo(enc *encoder) {
	enc.str(m.msg)
}

func (m *helloResponse) decodeFrom(dec *decoder) {
	m.msg = dec.str()
}
