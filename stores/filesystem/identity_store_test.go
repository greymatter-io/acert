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
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIdentityStore(t *testing.T) {

	Convey("IdentityStore", t, func() {

		Convey(".Fetch", func() {

			store := NewIdentityStore("./testdata")

			Convey("when the identity does not exist", func() {

				identity, err := store.Fetch("e3b0c44298fc")

				Convey("it returns a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})

				Convey("it returns a nil identity", func() {
					So(identity, ShouldBeNil)
				})
			})

			Convey("when the identity does exist", func() {

				identity, err := store.Fetch("6556CB34ADF5")

				Convey("it returns a nil error", func() {
					So(err, ShouldBeNil)
				})

				Convey("it returns the identity", func() {
					So(identity, ShouldNotBeNil)
				})
			})
		})

		Convey(".List", func() {

			Convey("when the directory is empty", func() {

				directory, err := ioutil.TempDir("", "empty")
				if err != nil {
					t.Fail()
				}

				store := NewIdentityStore(directory)
				list, err := store.List()

				Convey("it returns a nil error", func() {
					So(err, ShouldBeNil)
				})

				Convey("it returns a empty array", func() {
					So(list, ShouldBeEmpty)
				})
			})

			Convey("when the directory is not empty", func() {

				store := NewIdentityStore("./testdata")
				list, err := store.List()

				Convey("it returns a nil error", func() {
					So(err, ShouldBeNil)
				})

				Convey("it returns a non empty array", func() {
					So(list, ShouldNotBeEmpty)
				})
			})
		})
	})
}
