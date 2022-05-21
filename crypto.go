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
	"strconv"

	"golang.org/x/crypto/nacl/box"
)

// Encrypt msg with public key using nacl box seal anonymous.
//  - Parameter "base64PublicKey" and returning string are base64 encoded.
//  - Return nil if error
//  - All errors append to Errs
func BoxSealAnonymous(base64PublicKey, msg *string) *string {
	var err error
	var decodedKeyByte []byte
	// Check incoming base64 public key decode length. Must be 32 bytes.
	var length int = base64.StdEncoding.DecodedLen(len(*base64PublicKey))
	if length != 32 {
		err = Err("Decoded key length is " + strconv.Itoa(length) + ". Must be 32.")
	}
	// Decode incoming base64 public key
	if err == nil {
		decodedKeyByte = make([]byte, 32)
		_, err = base64.StdEncoding.Decode(decodedKeyByte, []byte(*base64PublicKey))
	}
	// Encrypt incoming message with public key
	var encryptedMsgByte []byte
	if err == nil {
		var publicKeyByte [32]byte
		copy(publicKeyByte[:], decodedKeyByte)
		encryptedMsgByte, err = box.SealAnonymous(nil, []byte(*msg), &publicKeyByte, nil)
	}
	// Encode the encrypted message to base64
	var base64EncryptedMsg string
	if err == nil {
		base64EncryptedMsg = base64.StdEncoding.EncodeToString(encryptedMsgByte)
	}
	// Add error to Errs
	if err != nil {
		Errs = append(Errs, err)
	}

	return &base64EncryptedMsg
}
