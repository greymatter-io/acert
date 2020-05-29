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
	"time"
)

// Options defines the options for the issue command.
type Options struct {

	// CommonName defines the common name for an certificate.
	CommonName string `mapstructure:"commonName"`

	// Country defines the country for an certificate.
	Country string `mapstructure:"country"`

	// DNSNames defines the subject alternative names for a certificate.
	DNSNames []string `mapstructure:"dnsNames"`

	// Expires defines the duration for which an certificate is valid.
	Expires time.Duration `mapstructure:"expires"`

	// Locality defines the city or county for an certificate.
	Locality string `mapstructure:"locality"`

	// Organization defines the organization for an certificate.
	Organization string `mapstructure:"organization"`

	// OrganizationalUnit defines the organization unit for an certificate.
	OrganizationalUnit string `mapstructure:"organizationalUnit"`

	// PostalCode defines the postal code for an certificate.
	PostalCode string `mapstructure:"postalCode"`

	// State defines the state or province for an certificate.
	State string `mapstructure:"state"`

	// StreetAddress defines the street address for an certificate.
	StreetAddress string `mapstructure:"streetAddress"`
}
