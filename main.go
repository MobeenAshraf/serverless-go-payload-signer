package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/keys"
	"github.com/coinbase/rosetta-sdk-go/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Result struct {
	SigningMessage string
	PrivKeyHex     string
	SenderAddress  string
	SigType        string
	PayloadSigType string
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var result Result
	json.Unmarshal([]byte(request.Body), &result)

	signingMessage := result.SigningMessage
	privKeyHex := result.PrivKeyHex
	senderAddress := result.SenderAddress

	if result.PayloadSigType == "" {
		result.PayloadSigType = "ecdsa_recovery"
	}
	if result.SigType == "" {
		result.SigType = "ecdsa_recovery"
	}

	var (
		payloadSigType types.SignatureType = types.SignatureType(result.PayloadSigType)
		sigType        types.SignatureType = types.SignatureType(result.SigType)
	)

	signingPayloadHexDecoded, _ := hex.DecodeString(signingMessage)
	keyPair, err := keys.ImportPrivateKey(privKeyHex, types.Secp256k1)
	if err != nil {
		return returnErr(err)
	}

	signer, _ := keyPair.Signer()
	sig, err := signer.Sign(&types.SigningPayload{
		AccountIdentifier: &types.AccountIdentifier{
			Address: senderAddress,
		},
		Bytes:         signingPayloadHexDecoded,
		SignatureType: payloadSigType,
	}, sigType)
	if err != nil {
		return returnErr(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       hex.EncodeToString(sig.Bytes),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}

func returnErr(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Body:       fmt.Sprint(err),
	}, nil
}
