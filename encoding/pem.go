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
	"encoding/pem"
)

// PEMEncodeCertificate returns the PEM encoded string for an X.509 certificate.
func PEMEncodeCertificate(certificate *x509.Certificate) string {
	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certificate.Raw}))
}

// PEMEncodeCertificates returns the PEM encoded strings for an array of X.509 certificates.
func PEMEncodeCertificates(certificates []*x509.Certificate) []string {

	pems := make([]string, len(certificates))
	for index, certificate := range certificates {
		pems[index] = string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certificate.Raw}))
	}

	return pems
}

// PEMEncodeKey returns the PEM encoded string for an RSA key.
func PEMEncodeKey(key *rsa.PrivateKey) string {
	return string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}))
}
