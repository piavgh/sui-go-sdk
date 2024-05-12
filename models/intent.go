package models

import (
	"github.com/block-vision/sui-go-sdk/constant"
)

type AppId int

const (
	Sui AppId = 0
)

type IntentVersion int

const (
	V0 IntentVersion = 0
)

func IntentWithScope(intentScope constant.IntentScope) []int {
	return []int{int(intentScope), int(V0), int(Sui)}
}

func MessageWithIntent(scope constant.IntentScope, message []byte) []byte {
	intent := []byte{uint8(scope), uint8(V0), uint8(Sui)}
	intentMessage := make([]byte, len(intent)+len(message))
	copy(intentMessage, intent)
	copy(intentMessage[len(intent):], message)
	return intentMessage
}
