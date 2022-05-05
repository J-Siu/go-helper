/*
Copyright © 2022 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package helper

import (
	"encoding/base64"

	"golang.org/x/crypto/nacl/box"
)

// Encrypt msg with public key using nacl box seal anonymous.
//
// Parameter "base64Pubkey" and returning string are base64 encoded.
func BoxSealAnonymous(base64Pubkey, msg *string) *string {
	keyByte := make([]byte, base64.StdEncoding.DecodedLen(len(*base64Pubkey)))
	base64.StdEncoding.Decode(keyByte, []byte(*base64Pubkey))
	keyByteP := new([32]byte)
	copy((*keyByteP)[:], keyByte)
	outByte, _ := box.SealAnonymous(nil, []byte(*msg), keyByteP, nil)
	outStr := base64.StdEncoding.EncodeToString(outByte)
	return &outStr
}
