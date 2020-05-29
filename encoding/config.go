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

package encoding

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"net/url"
	"strings"
)

// ConfigEncodeCertificate returns a X.509 certificate encoded for use in an IdentityConfig.
func ConfigEncodeCertificate(certificate *x509.Certificate) string {
	return url.PathEscape(base64.StdEncoding.EncodeToString([]byte(PEMEncodeCertificate(certificate))))
}

// ConfigEncodeCertificates returns X.509 certificates encoded for use in an IdentityConfig.
func ConfigEncodeCertificates(certificates []*x509.Certificate) string {

	pems := make([]string, len(certificates))
	for index, certificate := range certificates {
		pems[index] = PEMEncodeCertificate(certificate)
	}

	return url.PathEscape(base64.StdEncoding.EncodeToString([]byte(strings.Join(pems, "\n"))))
}

// ConfigEncodeKey returns an RSA key encoded for use in an IdentityConfig.
func ConfigEncodeKey(key *rsa.PrivateKey) string {
	return url.PathEscape(base64.StdEncoding.EncodeToString([]byte(PEMEncodeKey(key))))
}
