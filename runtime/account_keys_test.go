/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package runtime

import (
	"fmt"
	"github.com/onflow/cadence/runtime/interpreter"
	"github.com/onflow/cadence/runtime/sema"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/cadence/runtime/stdlib"
	"github.com/onflow/cadence/runtime/tests/utils"
)

func TestAccountKey(t *testing.T) {

	t.Parallel()

	runtime := NewInterpreterRuntime()

	script := []byte(`
		pub fun main(): AccountKey {
			let key = AccountKey(
				PublicKey2(
					publicKey: "0102".decodeHex(),
					signAlgo: "ECDSA_P256"
				),
				hashAlgo: "SHA3_256",
				weight: 1.7
			)

			return key
      	}
	`)

	runtimeInterface := &testRuntimeInterface{}

	result, err := runtime.ExecuteScript(
		Script{
			Source: script,
		},
		Context{
			Interface: runtimeInterface,
			Location:  utils.TestLocation,
		},
	)
	require.NoError(t, err)
	fmt.Println(result)
}

func TestAuthAccountAddPublicKey(t *testing.T) {
	t.Parallel()

	runtime := NewInterpreterRuntime()

	var tests = []TestCase{
		{
			name: "Single key",
			code: `
				transaction {
					prepare(signer: AuthAccount) {
						let key = PublicKey2(
							publicKey: "0102".decodeHex(),
							signAlgo: "ECDSA_P256"
						)

						signer.keys.add(
							publicKey: key,
							hashAlgo: "SHA3_256",
							weight: 100.0
						)
					}
				}`,
			keyCount: 1,
			args:     []cadence.Value{},
			expected: []*PublicKey{
				newPublicKeyValue([]byte{1, 2}, "ECDSA_P256"),
			},
		},
	}

	for _, test := range tests {

		storage := &Storage{
			events: make([]cadence.Event, 0),
			keys:   make([]*PublicKey, 0),
		}

		runtimeInterface := getRuntimeInterface(storage)

		t.Run(test.name, func(t *testing.T) {
			err := execute(test, runtime, runtimeInterface)

			require.NoError(t, err)
			assert.Len(t, storage.keys, test.keyCount)
			assert.Equal(t, test.expected, storage.keys)

			for _, event := range storage.events[0:] {
				assert.EqualValues(t, stdlib.AccountKeyAddedEventType.ID(), event.Type().ID())
			}
		})
	}
}

func TestAuthAccountAddPublicKey2(t *testing.T) {
	t.Parallel()

	runtime := NewInterpreterRuntime()

	keyA := newPublicKeyExportedValue([]byte{1, 2, 3}, "ECDSA_P256")
	keyB := newPublicKeyExportedValue([]byte{4, 5, 6}, "ECDSA_P256")
	keys := cadence.NewArray([]cadence.Value{keyA, keyB})

	var tests = []TestCase{
		{
			name: "Multiple keys",
			code: `
				transaction(keys: [PublicKey2]) {
					prepare(signer: AuthAccount) {
						let acct = AuthAccount(payer: signer)
						for key in keys {
							signer.keys.add(
								publicKey: key,
								hashAlgo: "SHA3_256",
								weight: 100.0
							)
						}
					}
				}
			`,
			keyCount: 2,
			args:     []cadence.Value{keys},
			expected: []*PublicKey{
				newPublicKeyValue([]byte{1, 2, 3}, "ECDSA_P256"),
				newPublicKeyValue([]byte{4, 5, 6}, "ECDSA_P256"),
			},
		},
	}

	for _, test := range tests {

		storage := &Storage{
			events: make([]cadence.Event, 0),
			keys:   make([]*PublicKey, 0),
		}

		runtimeInterface := getRuntimeInterface(storage)

		t.Run(test.name, func(t *testing.T) {
			err := execute(test, runtime, runtimeInterface)

			require.NoError(t, err)
			assert.Len(t, keys.Values, test.keyCount)
			assert.Equal(t, test.expected, storage.keys)

			assert.EqualValues(t, stdlib.AccountCreatedEventType.ID(), storage.events[0].Type().ID())

			for _, event := range storage.events[1:] {
				assert.EqualValues(t, stdlib.AccountKeyAddedEventType.ID(), event.Type().ID())
			}
		})
	}
}

func TestAuthAccountAddPublicKeyErrors(t *testing.T) {
	t.Parallel()

	runtime := NewInterpreterRuntime()

	var tests = []TestCase{
		{
			name: "AccountKey as transaction param",
			code: `
				transaction(keys: [AccountKey]) {
					prepare(signer: AuthAccount) {
					}
				}
			`,
			err: "transaction parameter must be storable: `[AccountKey]`",
		},
	}

	for _, test := range tests {

		storage := &Storage{
			events: make([]cadence.Event, 0),
			keys:   make([]*PublicKey, 0),
		}

		runtimeInterface := getRuntimeInterface(storage)

		t.Run(test.name, func(t *testing.T) {
			err := execute(test, runtime, runtimeInterface)
			require.Error(t, err)
			assert.Contains(t, err.Error(), test.err)
		})
	}
}

// Utility methods

func getRuntimeInterface(storage *Storage) *testRuntimeInterface {
	return &testRuntimeInterface{
		storage: newTestStorage(nil, nil),
		getSigningAccounts: func() ([]Address, error) {
			return []Address{{42}}, nil
		},
		createAccount: func(payer Address) (address Address, err error) {
			return Address{42}, nil
		},
		addAccountKey: func(address Address, publicKey *PublicKey) (*AccountKey, error) {
			storage.keys = append(storage.keys, publicKey)
			return &AccountKey{
				KeyIndex:  interpreter.NewIntValueFromInt64(int64(len(storage.keys))),
				PublicKey: publicKey,
				HashAlgo:  nil,
				Weight:    0,
				IsRevoked: false,
			}, nil
		},
		emitEvent: func(event cadence.Event) error {
			storage.events = append(storage.events, event)
			return nil
		},
		decodeArgument: func(b []byte, t cadence.Type) (value cadence.Value, err error) {
			return jsoncdc.Decode(b)
		},
	}
}

func execute(test TestCase, runtime Runtime, runtimeInterface *testRuntimeInterface) error {
	args := encodeArgs(test.args)
	err := runtime.ExecuteTransaction(
		Script{
			Source:    []byte(test.code),
			Arguments: args,
		},
		Context{
			Interface: runtimeInterface,
			Location:  utils.TestLocation,
		},
	)
	return err
}

func encodeArgs(argValues []cadence.Value) [][]byte {
	args := make([][]byte, len(argValues))
	for i, arg := range argValues {
		var err error
		args[i], err = jsoncdc.Encode(arg)
		if err != nil {
			panic(fmt.Errorf("broken test: invalid argument: %w", err))
		}
	}
	return args
}

func newPublicKeyValue(keyBytes []byte, signAlgo string) *PublicKey {
	intValues := make([]interpreter.Value, len(keyBytes))
	for index, value := range keyBytes {
		intValues[index] = interpreter.UInt8Value(value)
	}

	return &PublicKey{
		PublicKey: interpreter.NewArrayValueUnownedNonCopying(intValues...),
		SignAlgo:  interpreter.NewStringValue(signAlgo),
	}
}

func newPublicKeyExportedValue(keyBytes []byte, signAlgo string) cadence.BuiltinStruct {
	byteArray := make([]cadence.Value, len(keyBytes))
	for index, value := range keyBytes {
		byteArray[index] = cadence.NewUInt8(value)
	}

	return cadence.BuiltinStruct{
		StructType: PublicKeyType,
		Fields: []cadence.Value{
			cadence.NewArray(byteArray),
			cadence.NewString("ECDSA_P256"),
		},
	}
}

var PublicKeyType = func() *cadence.BuiltinStructType {
	var fields = []cadence.Field{
		{
			Identifier: sema.PublicKeyPublicKeyField,
			Type:       cadence.VariableSizedArrayType{ElementType: cadence.Int8Type{}},
		},
		{
			Identifier: sema.PublicKeySignAlgoField,
			Type:       cadence.StringType{},
		},
	}

	return &cadence.BuiltinStructType{
		QualifiedIdentifier: sema.PublicKeyTypeName,
		Fields:              fields,
	}
}()

type TestCase struct {
	name     string
	code     string
	keyCount int
	args     []cadence.Value
	expected []*PublicKey
	err      string
}

type Storage struct {
	events []cadence.Event
	keys   []*PublicKey
}
