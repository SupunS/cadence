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

// Contains all jump instructions.

// GOTO instruction
type GOTO struct {
	Instruction int
}

var _ vm.Instruction = GOTO{}

func (i GOTO) Execute(vm *vm.VirtualMachine) {
	vm.NextIndex = i.Instruction
}

func (i GOTO) String() string {
	return fmt.Sprintf("GOTO %d", i.Instruction)
}


// ICOMP instruction
type ICOMP struct {
	Instruction int // instruction to jump to, if false
}

var _ vm.Instruction = ICOMP{}

func (i ICOMP) Execute(vm *vm.VirtualMachine) {
	rhsOp := vm.CurrentStackFrame().Pop().(int)
	lhsOp := vm.CurrentStackFrame().Pop().(int)

	if lhsOp != rhsOp {
		vm.NextIndex = i.Instruction
	}
}

func (i ICOMP) String() string {
	return fmt.Sprintf("ICONST %d", i.Instruction)
}
