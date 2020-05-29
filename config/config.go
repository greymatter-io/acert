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

package config

import (
	"os"
	"path/filepath"

	"github.com/greymatter-io/acert/stores/filesystem"
	"github.com/pkg/errors"
)

// Authorities returns the authority identity store.
func Authorities() (*filesystem.IdentityStore, error) {

	directory, err := relative("authorities")
	if err != nil {
		return nil, errors.Wrap(err, "error determining authorities directory")
	}

	return filesystem.NewIdentityStore(directory), nil
}

// Leaves returns the leaf identity store.
func Leaves() (*filesystem.IdentityStore, error) {

	directory, err := relative("leaves")
	if err != nil {
		return nil, errors.Wrap(err, "error determining leaves directory")
	}

	return filesystem.NewIdentityStore(directory), nil
}

// relative returns the absolute path to a relative path in the configuration directory.
func relative(path string) (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "error determining user home directory")
	}

	return filepath.Join(home, ".acert", path), nil
}
