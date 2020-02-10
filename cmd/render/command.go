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

package render

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/greymatter-io/templar/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Command returns a command that renders one or more templates.
func Command() *cobra.Command {

	command := &cobra.Command{
		Use:   "render [target...]",
		Short: "Render one or more templates",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(command *cobra.Command, args []string) error {

			context := &templates.Context{Variables: make(map[string]string)}

			flags, err := command.Flags().GetStringArray("variable")
			if err != nil {
				return errors.Wrap(err, "error retrieving variable flags")
			}

			err = mergeCommandVariables(flags, context.Variables)
			if err != nil {
				return errors.Wrap(err, "error parsing variable flags")
			}

			flags, err = command.Flags().GetStringArray("environment")
			if err != nil {
				return errors.Wrap(err, "error retrieving environment flags")
			}

			err = mergeEnvironmentVariables(flags, context.Variables)
			if err != nil {
				return errors.Wrap(err, "error parsing variable flags")
			}

			targets, err := parseTargets(args)
			if err != nil {
				return errors.Wrap(err, "error parsing targets")
			}

			for _, target := range targets {
				err := templates.Render(target, context)
				if err != nil {
					return errors.Wrapf(err, "error rendering template from [%s] to [%s]", target.Source, target.Destination)
				}
			}

			return nil
		},
	}

	command.Flags().StringArrayP("variable", "v", []string{}, "the name and value of a user variable (i.e., NAME=VALUE)")
	command.Flags().StringArrayP("environment", "e", []string{}, "the name of an environment variable")

	return command
}

// mergeCommandVariable merges an unparsed command line variable into a variable map overwriting an existing entry if present.
func mergeCommandVariable(value string, variables map[string]string) error {
	split := strings.Split(value, "=")
	if len(split) != 2 {
		return fmt.Errorf("error parsing variable [%s], must be in the format NAME=VALUE", value)
	}
	variables[split[0]] = split[1]
	return nil
}

// mergeCommandVariables merges a slice of unparsed command line variables into a variable map overwriting an existing entries if present.
func mergeCommandVariables(values []string, variables map[string]string) error {
	for index, variable := range values {
		err := mergeCommandVariable(variable, variables)
		if err != nil {
			return errors.Wrapf(err, "error merging command variable at index [%d] and value [%s]", index, variable)
		}
	}
	return nil
}

// mergeEnvironmentVariable merges an environment variable into a variable map overwriting an existing entry if present.
func mergeEnvironmentVariable(value string, variables map[string]string) error {
	variable, found := os.LookupEnv(value)
	if !found {
		return fmt.Errorf("error processing unset environment variable [%s]", value)
	}
	variables[value] = variable
	return nil
}

// mergeEnvironmentVariables merges a slice of environment variables into a variable map overwriting an existing entries if present.
func mergeEnvironmentVariables(values []string, variables map[string]string) error {
	for index, variable := range values {
		err := mergeEnvironmentVariable(variable, variables)
		if err != nil {
			return errors.Wrapf(err, "error merging environment variable at index [%d] and value [%s]", index, variable)
		}
	}
	return nil
}

// parseTarget parses target argument into a target instance.
func parseTarget(value string) (*templates.Target, error) {
	split := strings.Split(value, ":")
	switch len(split) {
	case 2:
		return &templates.Target{Destination: split[1], Mode: 0400, Source: split[0]}, nil
	case 3:
		mode, err := strconv.ParseUint(split[2], 8, 32)
		if err != nil {
			return nil, errors.Wrapf(err, "error parsing mode of target [%s]", value)
		}
		return &templates.Target{Destination: split[1], Mode: os.FileMode(mode), Source: split[0]}, nil
	default:
		return nil, fmt.Errorf("error parsing target [%s]", value)
	}
}

// parseTargets parses a slice of target arguments into a slice of target instances.
func parseTargets(values []string) ([]*templates.Target, error) {
	targets := make([]*templates.Target, len(values))
	for index, value := range values {
		target, err := parseTarget(value)
		if err != nil {
			return nil, errors.Wrapf(err, "error parsing target at index [%d] with value [%s]", index, value)
		}
		targets[index] = target
	}
	return targets, nil
}
