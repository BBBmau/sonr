package simulation

import (
	"crypto/ed25519"
	cryptrand "crypto/rand"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/x/schema/types"
)

// FindAccount find a specific address from an account list
func FindAccount(accs []simtypes.Account, address string) (simtypes.Account, bool) {
	creator, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		panic(err)
	}
	return simtypes.FindAccount(accs, creator)
}

func CreatMockWhatIs(simmAcc simtypes.Account) (types.WhatIs, error) {
	schema, err := CreateMockSchema(simmAcc)
	if err != nil {
		return types.WhatIs{}, err
	}

	return types.WhatIs{
		Did:       schema.GetDid(),
		Schema:    &schema,
		Creator:   simmAcc.Address.String(),
		Timestamp: time.Now().Unix(),
		IsActive:  true,
	}, nil
}

func CreateMockSchema(simAcc simtypes.Account) (types.Schema, error) {
	doc, err := CreateMockDidDocument(simAcc)
	if err != nil {
		panic(err)
	}
	schema := types.Schema{
		Did:   doc.GetID().String(),
		Label: "test",
		Fields: []*types.SchemaField{
			{
				Name: "name",
				FieldKind: &types.SchemaFieldKind{
					Kind: types.Kind_STRING,
				},
			},
			{
				Name: "age",
				FieldKind: &types.SchemaFieldKind{
					Kind: types.Kind_INT,
				},
			},
			{
				Name: "DOB",
				FieldKind: &types.SchemaFieldKind{
					Kind: types.Kind_STRING,
				},
			},
			{
				Name: "categories",
				FieldKind: &types.SchemaFieldKind{
					Kind: types.Kind_LIST,
					ListKind: &types.SchemaFieldKind{
						Kind: types.Kind_STRING,
					},
				},
			},
		},
	}

	return schema, nil
}

// CreateMockDidDocument creates a mock did document for testing
func CreateMockDidDocument(simAccount simtypes.Account) (did.Document, error) {
	rawCreator := simAccount.Address.String()

	// Trim snr account prefix
	if strings.HasPrefix(rawCreator, "snr") {
		rawCreator = strings.TrimLeft(rawCreator, "snr")
	}

	// Trim cosmos account prefix
	if strings.HasPrefix(rawCreator, "cosmos") {
		rawCreator = strings.TrimLeft(rawCreator, "cosmos")
	}

	// UnmarshalJSON from DID document
	doc, err := did.NewDocument(fmt.Sprintf("did:snr:%s", rawCreator))
	if err != nil {
		return nil, err
	}

	//webauthncred := CreateMockCredential()
	pubKey, _, err := ed25519.GenerateKey(cryptrand.Reader)
	if err != nil {
		return nil, err
	}

	didUrl, err := did.ParseDID(fmt.Sprintf("did:snr:%s", rawCreator))
	if err != nil {
		return nil, err
	}
	didController, err := did.ParseDID(fmt.Sprintf("did:snr:%s#test", rawCreator))
	if err != nil {
		return nil, err
	}

	vm, err := did.NewVerificationMethod(*didUrl, ssi.JsonWebKey2020, *didController, pubKey)
	if err != nil {
		return nil, err
	}
	doc.AddAuthenticationMethod(vm)
	return doc, nil
}
