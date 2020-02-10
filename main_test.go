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

package main

import (
	"fmt"
	"testing"

	"github.com/greymatter-io/templar/internal/tests"
	. "github.com/smartystreets/goconvey/convey"
)

func TestVersion(t *testing.T) {

	Convey("When `templar version` is invoked", t, func() {

		commit := tests.MustGenerateHex(t)
		commitFlag := fmt.Sprintf("-X github.com/greymatter-io/templar/versions.commit=%s", commit)
		version := tests.MustGenerateHex(t)
		versionFlag := fmt.Sprintf("-X github.com/greymatter-io/templar/versions.version=%s", version)
		ldFlags := fmt.Sprintf("%s %s", versionFlag, commitFlag)

		stdout, stderr := tests.MustCommand(t, "go", "run", "-ldflags", ldFlags, "main.go", "version")

		Convey("it prints the version to stdout", func() {
			So(string(stdout), ShouldEqual, fmt.Sprintf("templar %s (%s)\n", version, commit))
		})

		Convey("it prints nothing to stderr", func() {
			So(string(stderr), ShouldEqual, "")
		})
	})
}
