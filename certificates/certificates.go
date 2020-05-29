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

package certificates

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"time"
)

// CommonName returns the common name of a certificate.
func CommonName(certificate *x509.Certificate) string {
	return certificate.Subject.CommonName
}

// Expiration returns the expiration of a certificate.
func Expiration(certificate *x509.Certificate) time.Time {
	return certificate.NotAfter
}

// Fingerprint returns the SHA256 hash of a certificate truncated to twelve characters.
func Fingerprint(certificate *x509.Certificate) string {
	bytes := sha256.Sum256(certificate.Raw)
	return hex.EncodeToString(bytes[:])[0:12]
}
