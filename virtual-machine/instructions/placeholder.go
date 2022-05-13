/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2022 Dapper Labs, Inc.
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

package instructions

import (
	"fmt"
	vm "github.com/onflow/cadence/virtual-machine"
)

// Function instruction is just a placeholder, to mark the start of a function.
// This is not needed in the future, just here for testing purpose.
type Function struct {
	Name string
}

func NewFunction(name string) Function {
	return Function{
		name,
	}
}

var _ vm.Instruction = Function{}

func (i Function) Execute(*vm.VirtualMachine) {
	//no-op
}

func (i Function) String() string {
	return fmt.Sprintf("FUNCTION %s:", i.Name)
}
