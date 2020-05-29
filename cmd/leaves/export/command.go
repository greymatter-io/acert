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

package export

import (
	"fmt"
	"strings"

	"github.com/greymatter-io/acert/config"
	"github.com/greymatter-io/acert/encoding"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Command returns a command that exports a leaf.
func Command() *cobra.Command {

	command := &cobra.Command{
		Use:   "export FINGERPRINT",
		Short: "Export a leaf",
		Args:  cobra.ExactArgs(1),
		RunE: func(command *cobra.Command, args []string) error {

			viper.BindPFlag("format", command.Flags().Lookup("format"))
			viper.BindPFlag("type", command.Flags().Lookup("type"))

			var options Options

			err := viper.Unmarshal(&options)
			if err != nil {
				return err
			}

			leaves, err := config.Leaves()
			if err != nil {
				return err
			}

			leaf, err := leaves.Fetch(args[0])
			if err != nil {
				return err
			}

			switch strings.ToLower(options.Type) {
			case "authority":
				fmt.Println(strings.Join(encoding.PEMEncodeCertificates(leaf.Authorities), ""))
				break
			case "certificate":
				fmt.Println(encoding.PEMEncodeCertificate(leaf.Certificate))
				break
			case "key":
				fmt.Println(encoding.PEMEncodeKey(leaf.Key))
				break
			default:
				return fmt.Errorf("error parsing type [%s] must be one of [authority, certificate, key]", options.Type)
			}

			return nil
		},
	}

	command.Flags().StringP("format", "f", "pem", "the format of the exported leaf")
	command.Flags().StringP("type", "t", "certificate", "the type of values to be exported [authority, certificate, key]")

	return command
}
