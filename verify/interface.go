package verify

import "github.com/machinebox/graphql"

type IPublicKey interface {
	VerifyPersonalMessage(message []byte, signature []byte, client *graphql.Client) (string, bool, error)
}
