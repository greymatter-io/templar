// Copyright 2020 Decipher Technology Studios
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"io"
	"io/ioutil"
	"testing"
)

// MustReadAndClose reads initializes a command and returns the command and pipes for stdin, stdout and stderr or fails the test.
func MustReadAndClose(t *testing.T, r io.ReadCloser) []byte {

	defer r.Close()

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error("unable to read reader")
	}

	return bytes
}
