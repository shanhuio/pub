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

package bosinit

import (
	"testing"

	"reflect"
	"strings"
)

func TestParseConfig(t *testing.T) {
	content := strings.Join([]string{
		"ssh_authorized_keys:",
		"- ssh-rsa key1",
		"- ssh-rsa key2",
	}, "\n")

	config, err := ParseConfig([]byte(content))
	if err != nil {
		t.Fatal("parse config: ", err)
	}

	wantKeys := []string{
		"ssh-rsa key1",
		"ssh-rsa key2",
	}
	if got := config.SSHAuthorizedKeys; !reflect.DeepEqual(got, wantKeys) {
		t.Errorf("parse config, got keys %q, want %q", got, wantKeys)
	}
}
