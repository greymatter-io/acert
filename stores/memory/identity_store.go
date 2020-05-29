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

package memory

import (
	"fmt"

	"github.com/greymatter-io/acert/certificates"
	"github.com/greymatter-io/nautls/identities"
)

// IdentityStore provides an in memory implementation of the IdentityStore interface.
type IdentityStore struct {
	store map[string]*identities.Identity
}

// Delete deletes the identity with the provided fingerprint from this store.
func (s *IdentityStore) Delete(fingerprint string) error {
	delete(s.store, fingerprint)
	return nil
}

// Fetch returns the identity with the provided fingerprint from this store.
func (s *IdentityStore) Fetch(fingerprint string) (*identities.Identity, error) {

	identity, found := s.store[fingerprint]
	if !found {
		return nil, fmt.Errorf("identity not found [%s]", fingerprint)
	}

	return identity, nil
}

// List returns the identities from this store.
func (s *IdentityStore) List() ([]*identities.Identity, error) {

	result := []*identities.Identity{}
	for _, value := range s.store {
		result = append(result, value)
	}

	return result, nil
}

// Upsert inserts or updates an identity into this store and returns the id.
func (s *IdentityStore) Upsert(identity *identities.Identity) (string, error) {
	fingerprint := certificates.Fingerprint(identity.Certificate)
	s.store[fingerprint] = identity
	return fingerprint, nil
}
