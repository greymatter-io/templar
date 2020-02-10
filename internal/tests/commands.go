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
	"os/exec"
	"testing"
)

// MustExecute executes a command and returns the values of stdout and stderr or fails the test.
func MustCommand(t *testing.T, name string, args ...string) ([]byte, []byte) {

	command := exec.Command(name, args...)

	stdoutPipe, err := command.StdoutPipe()
	if err != nil {
		t.Error("error opening stdout pipe")
	}

	stderrPipe, err := command.StderrPipe()
	if err != nil {
		t.Error("error opening stderr pipe")
	}

	err = command.Start()
	if err != nil {
		t.Error("error starting the command")
	}

	stdout := MustReadAndClose(t, stdoutPipe)
	stderr := MustReadAndClose(t, stderrPipe)

	err = command.Wait()
	if err != nil {
		t.Error("error waiting on command")
	}

	return stdout, stderr
}
