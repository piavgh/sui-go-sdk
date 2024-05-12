package models

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/blake2b"

	"github.com/block-vision/sui-go-sdk/constant"
)

func verifyWithIntent(message, signature string, scope constant.IntentScope) (signer string, pass bool, err error) {
	b64Bytes, _ := base64.StdEncoding.DecodeString(message)
	bcsBytes := append([]byte{uint8(len(b64Bytes))}, b64Bytes...)

	messageBytes := MessageWithIntent(scope, bcsBytes)

	serializedSignature, err := parseSerializedSignature(signature)
	if err != nil {
		return "", false, err
	}

	digest := blake2b.Sum256(messageBytes)
	pass = ed25519.Verify(serializedSignature.PubKey[:], digest[:], serializedSignature.Signature)
	signer = ed25519PublicKeyToSuiAddress(serializedSignature.PubKey)

	return
}

func VerifyPersonalMessage(message string, signature string) (signer string, pass bool, err error) {
	b64Message := base64.StdEncoding.EncodeToString([]byte(message))
	return verifyWithIntent(b64Message, signature, constant.PersonalMessage)
}

func VerifyTransaction(b64Message string, signature string) (signer string, pass bool, err error) {
	return verifyWithIntent(b64Message, signature, constant.TransactionData)
}

func ed25519PublicKeyToSuiAddress(pubKey []byte) string {
	newPubkey := []byte{byte(SigFlagEd25519)}
	newPubkey = append(newPubkey, pubKey...)

	addrBytes := blake2b.Sum256(newPubkey)
	return fmt.Sprintf("0x%s", hex.EncodeToString(addrBytes[:])[:64])
}
