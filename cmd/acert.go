// Copyright 2019 Decipher Technology Studios
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
	"github.com/greymatter-io/acert/cmd/authorities"
	"github.com/greymatter-io/acert/cmd/leaves"
	"github.com/greymatter-io/acert/cmd/version"
	"github.com/spf13/cobra"
)

// Acert returns a command that creates and manages X.509 identites.
func Acert() *cobra.Command {

	command := &cobra.Command{
		Use:   "acert",
		Short: "Manage X.509 identities",
		Long:  "A command line utility for creating and managing X.509 identities.",
	}

	command.AddCommand(authorities.Command())
	command.AddCommand(leaves.Command())
	command.AddCommand(version.Command())

	return command
}
