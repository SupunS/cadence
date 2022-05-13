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

// Store instruction reads a values from the top of the stack and
// stores it in a local variable at a given index.
type Store struct {
	Index int
}

var _ vm.Instruction = Store{}

func NewIStore(index int) Store {
	return Store{
		Index: index,
	}
}

func (i Store) Execute(vm *vm.VirtualMachine) {
	value := vm.CurrentStackFrame().Pop()
	vm.CurrentStackFrame().Set(i.Index, value)
}

func (i Store) String() string {
	return fmt.Sprintf("STORE %d", i.Index)
}
