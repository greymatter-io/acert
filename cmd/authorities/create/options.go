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

package create

import (
	"time"
)

// Options defines the options for the create command.
type Options struct {

	// CommonName defines the common name for an authority.
	CommonName string `mapstructure:"commonName"`

	// Country defines the country for an authority.
	Country string `mapstructure:"country"`

	// Expires defines the duration for which an authority is valid.
	Expires time.Duration `mapstructure:"expires"`

	// Locality defines the city or county for an authority.
	Locality string `mapstructure:"locality"`

	// Organization defines the organization for an authority.
	Organization string `mapstructure:"organization"`

	// OrganizationalUnit defines the organization unit for an authority.
	OrganizationalUnit string `mapstructure:"organizationalUnit"`

	// PostalCode defines the postal code for an authority.
	PostalCode string `mapstructure:"postalCode"`

	// State defines the state or province for an authority.
	State string `mapstructure:"state"`

	// StreetAddress defines the street address for an authority.
	StreetAddress string `mapstructure:"streetAddress"`
}
