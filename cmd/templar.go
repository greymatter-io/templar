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

package cmd

import (
	"github.com/greymatter-io/templar/cmd/render"
	"github.com/greymatter-io/templar/cmd/version"
	"github.com/spf13/cobra"
)

// Templar returns a command renders Go templates.
func Templar() *cobra.Command {

	command := &cobra.Command{
		Use:   "templar",
		Short: "Render Go templates.",
		Long:  "A command line utility for rendering Go templates to the filesystem..",
	}

	command.AddCommand(render.Command())
	command.AddCommand(version.Command())

	return command
}
