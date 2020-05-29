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

package stores

import "github.com/greymatter-io/nautls/identities"

// IdentityStore defines the interface for identity stores.
type IdentityStore interface {

	// Delete deletes the identity with the provided fingerprint from this store.
	Delete(fingerprint string) error

	// Fetch returns the identity with the provided fingerprint from this store.
	Fetch(fingerprint string) (*identities.Identity, error)

	// List returns the identities from this store.
	List() ([]*identities.Identity, error)

	// Upsert inserts or updates an identity into this store and returns the fingerprint.
	Upsert(*identities.Identity) (string, error)
}
