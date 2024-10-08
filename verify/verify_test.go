package verify

import (
	"testing"

	"github.com/machinebox/graphql"

	"github.com/block-vision/sui-go-sdk/mystenbcs"
	"github.com/block-vision/sui-go-sdk/zklogin"
)

func TestVerifyPersonalMessageSignature(t *testing.T) {
	graphqlClient := graphql.NewClient("https://sui-testnet.mystenlabs.com/graphql")

	type args struct {
		message   string
		signature string
		options   *zklogin.ZkLoginPublicIdentifierOptions
	}
	tests := []struct {
		name       string
		args       args
		wantSigner string
		wantPass   bool
		wantErr    bool
	}{
		{
			name: "Test VerifyPersonalMessageSignature",
			args: args{
				message:   "0x4c8e92e12f78ec6ce36fc07870baf8e8ffb313543b6398d92dc574af4a155987-GqwdQmUc9VFtmG9q",
				signature: "BQNNMTUzMTU1MTI4MzAwNjA0NTE4MjM2MTEzMTA5NjM1NjcyMTEzNTU2MTc4NjE0Mjc0NTMxNTg5MzY1MTc2OTE1NzM0NDQ0NjQ0NTA5ODdMMzgzNzI2Mjc4NzIzMDkzMDY2NzI1NTc2NTYwOTM0OTA4MTczNTUwMjYwNTg0NjAxMzYxNjEyNDk5ODE1ODE4OTA4ODQwOTE4NzM1OQExAwJNMTUwMDE3NzgyNDYxMzE2NzgyMjkyODk2NzA0OTg1NDkyNzcwOTIyMjk0NDAwNjI4NTM3OTg0ODA4MDcyNTY3NTAwODYxNTE5MDk3MTRLMzQzMDI0MDMzNTk3ODM1OTA0MzIxOTY1NDA0MDUyMDc3NTM3MzE3Mzc0MTU4ODkyODUzNTM5NTg5NzI0MDYyNTcwMjQ2MjkwODcyAk0xMDYzOTE2MTQzNTIxNDAxMTIxNzY3MzI0MzAxMjA2ODYxOTY4NzM4MDUzODAxMzQzNzQxNDczOTI4NDMzNzI3ODUxMTMxMTM4ODkwOU0xOTAyNzU5MzQwNTMzMzA3OTAyMzU2NDQxNjM5OTUwNjQ2NTgwNjk2MTk2MzY4NTY2MjAxMjQ2NTY4MTI5MDk2NzY3MzEzNzk5MTk1NAIBMQEwA0wzOTcxNzQ1ODI3NDI5OTQxNTA2OTUxMDE2OTA3NTc0ODY3NzMzODc0MjY1NzQ1MDQ0ODc5ODE3NjUyMjMxNzg4NDk5Mzk3ODEyNzY0TDc3NDA0MTcyMzA4MjMwNjg0NDEwNDc1NTgzNDQ5MzU4MDA1MDA1MzM0MDgwNjQ5NDEwMjk3OTkwNzU4MTU1MDU1ODU0NjUzNzUwMDIBMTF5SnBjM01pT2lKb2RIUndjem92TDJGalkyOTFiblJ6TG1kdmIyZHNaUzVqYjIwaUxDAWZleUpoYkdjaU9pSlNVekkxTmlJc0ltdHBaQ0k2SW1FMU1HWTJaVGN3WldZMFlqVTBPR0UxWm1RNU1UUXlaV1ZqWkRGbVlqaG1OVFJrWTJVNVpXVWlMQ0owZVhBaU9pSktWMVFpZlFMNDM5MDMzNzMyNzMzOTgwMjQ3MDc5ODMxNTY2MDExOTkyMjg4ODY3NjQwNTQzNjczNDU4MjczNzk3ODYzMTc3OTg5MzQ4OTIzMzQyMCICAAAAAAAAYQDiYfUblsMnZuYlq9KaZuJ3SfgcsvQ/GSY/wwvVeV7JXqMXCHPIArITmmZtxMRcmiGAE/9KTFSasIftYxHLk+cCopQm1DnsH/yOMHoIKuFQf2gi+Yg0NPobPh/a5cTor54=",
				options: &zklogin.ZkLoginPublicIdentifierOptions{
					Client: graphqlClient,
				},
			},
			wantSigner: "0x4c8e92e12f78ec6ce36fc07870baf8e8ffb313543b6398d92dc574af4a155987",
			wantPass:   true,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messageBytes, _ := mystenbcs.FromBase64(tt.args.message)
			signatureBytes, _ := mystenbcs.FromBase64(tt.args.signature)

			gotSigner, gotPass, err := VerifyPersonalMessageSignature(messageBytes, signatureBytes, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyPersonalMessageSignature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSigner != tt.wantSigner {
				t.Errorf("VerifyPersonalMessageSignature() gotSigner = %v, want %v", gotSigner, tt.wantSigner)
			}
			if gotPass != tt.wantPass {
				t.Errorf("VerifyPersonalMessageSignature() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}
