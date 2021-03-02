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

package sema

// PublicAccountType represents the publicly accessible portion of an account.
//
var PublicAccountType = func() *BuiltinCompositeType {

	publicAccountType := &BuiltinCompositeType{
		Identifier:           "PublicAccount",
		IsInvalid:            false,
		IsResource:           false,
		Storable:             false,
		Equatable:            false,
		ExternallyReturnable: false,

		nestedTypes: func() *StringTypeOrderedMap {
			nestedTypes := NewStringTypeOrderedMap()
			nestedTypes.Set(AccountKeysTypeName, PublicAccountKeysType)
			return nestedTypes
		}(),
	}

	var members = []*Member{
		NewPublicConstantFieldMember(
			publicAccountType,
			"address",
			&AddressType{},
			accountTypeAddressFieldDocString,
		),
		NewPublicConstantFieldMember(
			publicAccountType,
			"storageUsed",
			&UInt64Type{},
			accountTypeStorageUsedFieldDocString,
		),
		NewPublicConstantFieldMember(
			publicAccountType,
			"storageCapacity",
			&UInt64Type{},
			accountTypeStorageCapacityFieldDocString,
		),
		NewPublicFunctionMember(
			publicAccountType,
			"getCapability",
			authAccountTypeGetCapabilityFunctionType,
			authAccountTypeGetCapabilityFunctionDocString,
		),
		NewPublicFunctionMember(
			publicAccountType,
			"getLinkTarget",
			accountTypeGetLinkTargetFunctionType,
			publicAccountTypeGetLinkTargetFunctionDocString,
		),
		NewPublicConstantFieldMember(
			publicAccountType,
			"keys",
			PublicAccountKeysType,
			accountTypeKeysFieldDocString,
		),
	}

	publicAccountType.Members = GetMembersAsMap(members)
	return publicAccountType
}()

// PublicAccountKeysType represents the keys associated with a public account.
var PublicAccountKeysType = func() *BuiltinCompositeType {

	accountKeys := &BuiltinCompositeType{
		Identifier:           AccountKeysTypeName,
		IsInvalid:            false,
		IsResource:           false,
		Storable:             false,
		Equatable:            true,
		ExternallyReturnable: false,
	}

	var members = []*Member{
		NewPublicFunctionMember(
			accountKeys,
			AccountKeysGetFunctionName,
			accountKeysTypeGetFunctionType,
			accountKeysTypeGetFunctionDocString,
		),
	}

	accountKeys.Members = GetMembersAsMap(members)
	return accountKeys
}()

func init() {
	// Set the container type after initializing the AccountKeysTypes, to avoid initializing loop.
	PublicAccountKeysType.ContainerType = PublicAccountType
}

const publicAccountTypeGetLinkTargetFunctionDocString = `
Returns the capability at the given public path, or nil if it does not exist
`
