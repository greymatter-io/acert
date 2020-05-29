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

package filesystem

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/greymatter-io/acert/certificates"
	"github.com/greymatter-io/acert/encoding"
	"github.com/greymatter-io/nautls/identities"
	"github.com/pkg/errors"
)

// IdentityStore provides an on disk implementation of the IdentityStore interface.
type IdentityStore struct {
	directory string
}

// NewIdentityStore returns a new identity store instance.
func NewIdentityStore(directory string) *IdentityStore {
	return &IdentityStore{
		directory: directory,
	}
}

// Delete deletes the identity with the provided fingerprint from this store.
func (s *IdentityStore) Delete(fingerprint string) error {

	file := filepath.Join(s.directory, fmt.Sprintf("%s.json", fingerprint))

	err := os.Remove(file)
	if err != nil {
		return errors.Wrapf(err, "error deleting identity from [%s]", file)
	}

	return nil
}

// Fetch returns the identity with the provided fingerprint from this store.
func (s *IdentityStore) Fetch(fingerprint string) (*identities.Identity, error) {

	file := filepath.Join(s.directory, fmt.Sprintf("%s.json", fingerprint))

	identity, err := readIdentity(file)
	if err != nil {
		return nil, errors.Wrapf(err, "error loading identity from [%s]", file)
	}

	return identity, nil
}

// List returns an array of identities from this store.
func (s *IdentityStore) List() ([]*identities.Identity, error) {

	files, err := filepath.Glob(filepath.Join(s.directory, "*.json"))
	if err != nil {
		return nil, errors.Wrapf(err, "error reading contents of [%s]", s.directory)
	}

	identities := make([]*identities.Identity, len(files))

	for index, file := range files {

		identity, err := readIdentity(file)
		if err != nil {
			return nil, errors.Wrapf(err, "error loading identity from [%s]", file)
		}

		identities[index] = identity
	}

	return identities, nil
}

// Upsert inserts or updates an identity into this store and returns the fingerprint.
func (s *IdentityStore) Upsert(identity *identities.Identity) (string, error) {

	fingerprint := certificates.Fingerprint(identity.Certificate)

	err := writeIdentity(filepath.Join(s.directory, fmt.Sprintf("%s.json", fingerprint)), identity)
	if err != nil {
		return "", errors.Wrap(err, "error writing identity")
	}

	return fingerprint, nil
}

// readIdentity reads an identity from a file.
func readIdentity(path string) (*identities.Identity, error) {

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading identity from [%s]", path)
	}

	var config identities.IdentityConfig

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, errors.Wrapf(err, "error unmarshalling identity from [%s]", path)
	}

	identity, err := config.Build()
	if err != nil {
		return nil, errors.Wrapf(err, "error building identity from [%s]", path)
	}

	return identity, nil
}

// writeIdentity writes an identity to a file.
func writeIdentity(path string, identity *identities.Identity) error {

	config := &identities.IdentityConfig{
		Authorities: fmt.Sprintf("base64:///%s", encoding.ConfigEncodeCertificates(identity.Authorities)),
		Certificate: fmt.Sprintf("base64:///%s", encoding.ConfigEncodeCertificate(identity.Certificate)),
		Key:         fmt.Sprintf("base64:///%s", encoding.ConfigEncodeKey(identity.Key)),
	}

	bytes, err := json.Marshal(config)
	if err != nil {
		return errors.Wrapf(err, "error marshalling identity to file [%s]", path)
	}

	err = os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return errors.Wrapf(err, "error creating parent directory [%s]", filepath.Dir(path))
	}

	err = ioutil.WriteFile(path, bytes, 0600)
	if err != nil {
		return errors.Wrapf(err, "error writing identity to file [%s]", path)
	}

	return nil
}
