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

package delete

import (
	"github.com/greymatter-io/acert/config"
	"github.com/spf13/cobra"
)

// Command returns a command that deletes an authority.
func Command() *cobra.Command {

	command := &cobra.Command{
		Use:   "delete FINGERPRINT",
		Short: "Delete an authority",
		Args:  cobra.ExactArgs(1),
		RunE: func(command *cobra.Command, args []string) error {

			authorities, err := config.Authorities()
			if err != nil {
				return err
			}

			err = authorities.Delete(args[0])
			if err != nil {
				return err
			}

			return nil
		},
	}

	return command
}
