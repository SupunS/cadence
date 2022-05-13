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

// Load instruction reads a local variable from the given index
// of the stack and pushes it onto the top of the stack.
type Load struct {
	Index int
}

var _ vm.Instruction = Load{}

func NewLoad(index int) Load {
	return Load{
		Index: index,
	}
}

func (i Load) Execute(vm *vm.VirtualMachine) {
	value := vm.CurrentStackFrame().Get(i.Index)
	vm.CurrentStackFrame().Push(value)
}

func (i Load) String() string {
	return fmt.Sprintf("LOAD %d", i.Index)
}
