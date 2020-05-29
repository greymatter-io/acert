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

package issue

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"time"

	"github.com/greymatter-io/acert/config"
	"github.com/greymatter-io/nautls/identities"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Command returns a command that creates an authority.
func Command() *cobra.Command {

	command := &cobra.Command{
		Use:   "issue",
		Short: "Issue a certificate",
		Args:  cobra.ExactArgs(1),
		RunE: func(command *cobra.Command, args []string) error {

			viper.BindPFlag("commonName", command.Flags().Lookup("commonName"))
			viper.BindPFlag("country", command.Flags().Lookup("country"))
			viper.BindPFlag("expires", command.Flags().Lookup("expires"))
			viper.BindPFlag("state", command.Flags().Lookup("state"))
			viper.BindPFlag("locality", command.Flags().Lookup("locality"))
			viper.BindPFlag("organization", command.Flags().Lookup("organization"))
			viper.BindPFlag("organizationalUnit", command.Flags().Lookup("organizationalUnit"))
			viper.BindPFlag("postalCode", command.Flags().Lookup("postalCode"))
			viper.BindPFlag("streetAddress", command.Flags().Lookup("streetAddress"))

			var options Options

			err := viper.Unmarshal(&options)
			if err != nil {
				return err
			}

			template := identities.Template{
				BasicConstraintsValid: true,
				DNSNames:              options.DNSNames,
				ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
				KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
				NotAfter:              time.Now().Add(options.Expires),
				NotBefore:             time.Now(),
				SerialNumber:          big.NewInt(time.Now().Unix()),
				Subject: pkix.Name{
					CommonName:         options.CommonName,
					Country:            []string{options.Country},
					Locality:           []string{options.Locality},
					Organization:       []string{options.Organization},
					OrganizationalUnit: []string{options.OrganizationalUnit},
					Province:           []string{options.State},
				},
			}

			if options.StreetAddress != "" {
				template.Subject.StreetAddress = []string{options.StreetAddress}
			}

			if options.PostalCode != "" {
				template.Subject.PostalCode = []string{options.PostalCode}
			}

			authorities, err := config.Authorities()
			if err != nil {
				return err
			}

			authority, err := authorities.Fetch(args[0])
			if err != nil {
				return err
			}

			leaf, err := authority.Issue(template)
			if err != nil {
				return err
			}

			leaves, err := config.Leaves()
			if err != nil {
				return err
			}

			fingerprint, err := leaves.Upsert(leaf)
			if err != nil {
				return err
			}

			fmt.Println(fingerprint)

			return nil
		},
	}

	command.Flags().StringP("commonName", "n", "Acert", "common name for the authority")
	command.Flags().StringP("country", "c", "US", "two letter country code for the authority")
	command.Flags().DurationP("expires", "e", (time.Hour * 24 * 3650), "expiration time for the authority")
	command.Flags().StringP("state", "s", "Virginia", "state for the authority")
	command.Flags().StringP("locality", "l", "Alexandria", "locality for the authority")
	command.Flags().StringP("organization", "o", "Decipher Technology Studios", "organization for the authority")
	command.Flags().StringP("organizationalUnit", "u", "Engineering", "organizational unit for the authority")
	command.Flags().StringP("postalCode", "p", "", "postal code for the authority")
	command.Flags().StringP("streetAddress", "a", "", "street address for the authority")

	return command
}
